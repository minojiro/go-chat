package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Room struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type IRoom struct {
	Name string
}

var rooms = []Room{}

const timeLayout = "2006-01-02 15:04:05"

func main() {
	http.HandleFunc("/api/rooms", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			var data IRoom
			decoder.Decode(&data)
			room := Room{ID: len(rooms), Name: data.Name, CreatedAt: time.Now().Format(timeLayout)}
			fmt.Println(room)
			rooms = append(rooms, room)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rooms)
	})

	http.ListenAndServe(":8081", nil)
}
