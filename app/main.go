package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type (
	// Room is
	Room struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		CreatedAt   string `json:"created_at"`
		Connections []*websocket.Conn
	}

	// IRoom is
	IRoom struct {
		Name string `json:"name"`
	}

	// IMessage is
	IMessage struct {
		Body      string `json:"body"`
		CreatedBy string `json:"createdBy"`
		CreatedAt string `json:"createdAt"`
	}

	// InData is
	InData struct {
		DataType string `json:"dataType"`
		RoomID   int    `json:"roomId"`
		JoinData struct {
			Nickname string `json:"nickname"`
		} `json:"joinData"`
		MessageData struct {
			Body string `json:"body"`
		} `json:"messageData"`
	}
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	rooms = []Room{}
)

func getTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func main() {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	http.HandleFunc("/api/rooms", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			var data IRoom
			decoder.Decode(&data)
			room := Room{ID: len(rooms), Name: data.Name, CreatedAt: getTimeString()}
			fmt.Println(room)
			rooms = append(rooms, room)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rooms)
	})

	http.HandleFunc("/api/socket", func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("conn")

		defer conn.Close()

		var nickname string

		for {
			var d InData
			err := conn.ReadJSON(&d)
			if err != nil {
				fmt.Println(err)
				fmt.Println("CLOSE")
				break
			}

			room := &rooms[d.RoomID]

			if d.DataType == "join" {
				nickname = d.JoinData.Nickname
				room.Connections = append(room.Connections, conn)
			}

			if d.DataType == "message" {
				for _, otherConn := range room.Connections {
					message := IMessage{
						Body:      d.MessageData.Body,
						CreatedBy: nickname,
						CreatedAt: getTimeString(),
					}
					otherConn.WriteJSON(message)
				}
			}
			fmt.Println(d)
		}
	})

	http.ListenAndServe(":8081", nil)
}
