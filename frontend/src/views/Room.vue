<template>
  <div class="Room" v-if="room">
    <div class="chat" v-if="connection">
      <form @submit.prevent="createMessage">
        <div class="columns is-mobile">
          <div class="column is-four-fifths"><input :placeholder="`${nickname} says...`" class="input" type="text" v-model="newMessageBody"></div>
          <div class="column is-one-fifth"><button class="button is-primary is-fullwidth" :disabled="!newMessageBody">send</button></div>
        </div>
      </form>

      <br>

      <table class="table is-fullwidth">
        <tr v-for="message in messages" :key="message.id">
          <td>
            <div class="columns is-vcentered">
              <p class="column is-one-fifth">{{message.createdBy}}</p>
              <p class="column is-three-fifths">{{message.body}}</p>
              <p class="column has-text-right has-text-grey-light is-size-7">{{message.createdAt | datetime}}</p>
            </div>
          </td>
        </tr>
      </table>
    </div>

    <div class="join" v-else>
      <h2 class="title is-2">Type your nickname and join to {{room.name}}!</h2>
      <form @submit.prevent="join">
        <div class="columns is-mobile">
          <div class="column is-four-fifths"><input placeholder="nickname" class="input" type="text" v-model="nickname"></div>
          <div class="column is-one-fifth"><button class="button is-primary is-fullwidth" :disabled="!nickname">join</button></div>
        </div>
      </form>
    </div>
    <br>
    <br>
    <p><router-link to="/">back</router-link></p>
  </div>
</template>

<script>
import moment from 'moment'
import axios from 'axios'
import SocketIO from 'socket.io-client'

export default {
  name: 'Room',
  filters: {
    datetime(s) {
      const m = moment(s)
      return m.format('HH:mm:ss')
    },
  },
  data() {
    return {
      newMessageBody: '',
      nickname: '',
      room: {},
      connection: null,
      messages: [],
    }
  },
  props: {
    roomId: String,
  },
  methods: {
    createMessage() {
      const payload = {
        dataType: 'message',
        roomId: ~~this.roomId,
        messageData: {
          body: this.newMessageBody,
        },
      }
      this.connection.send(JSON.stringify(payload))
      this.newMessageBody = ''
    },
    join() {
      const connection = new WebSocket('ws://localhost:8081/api/socket')
      connection.onopen = (e) => {
        const payload = {
          dataType: 'join',
          roomId: ~~this.roomId,
          joinData: {
            nickname: this.nickname,
          },
        }
        connection.send(JSON.stringify(payload))
        connection.onmessage = this.recieveMessage
        this.connection = connection
      }
    },
    recieveMessage(e) {
      const message = JSON.parse(e.data)
      this.messages = [...this.messages, message].sort((a,b) => a.createdAt < b.createdAt ? 1 : -1).slice(0, 20)
    },
    async initialize() {
      const res = await axios.get('http://localhost:8081/api/rooms')
      this.room = res.data[~~this.roomId]
    },
  },
  async mounted() {
    this.initialize()
  },
  beforeDestroy() {
    if (this.connection) {
      this.connection.close()
    }
  },
}
</script>
