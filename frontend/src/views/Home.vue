<template>
  <div class="home">
    <h2 class="title is-1">create room</h2>
    <form @submit.prevent="createRoom">
      <div class="columns is-mobile">
        <div class="column is-four-fifths"><input placeholder="message" class="input" type="text" v-model="newRoomName"></div>
        <div class="column is-one-fifth"><button class="button is-primary is-fullwidth" :disabled="!newRoomName">create</button></div>
      </div>
    </form>

    <br><br>
    <div v-if="rooms.length">
      <h2 class="title is-1">rooms</h2>
      <table class="table is-fullwidth">
        <tr v-for="room in rooms" :key="room.id">
          <td>
            <div class="columns">
              <div class="column is-one-fifth">{{room.created_at | datetime}}</div>
              <div class="column is-four-fifths"><router-link :to="`/rooms/${room.id}`">{{room.name}}</router-link></div>
            </div>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import moment from 'moment'
import axios from 'axios'

export default {
  name: 'Home',
  filters: {
    datetime(s) {
      const m = new moment(s)
      return m.format('HH:mm:ss')
    },
  },
  data() {
    return {
      newRoomName: '',
      rooms: [],
    }
  },
  methods: {
    async fetchRooms() {
      const res = await axios.get('http://localhost:8081/api/rooms')
      this.rooms = res.data.sort((a,b) => a.createdAt < b.createdAt ? 1 : -1).slice(0, 10)
    },
    async createRoom() {
      const room = {
        name: this.newRoomName,
      }
      const res = await axios.post('http://localhost:8081/api/rooms', room)
      this.rooms = res.data.sort((a,b) => a.createdAt < b.createdAt ? 1 : -1).slice(0, 10)
    },
  },
  async mounted() {
    this.fetchRooms()
  },
}
</script>
