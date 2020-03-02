<template>
  <div class="Room" v-if="room">
    <h1 class="title is-1">{{room.name}}</h1>

    <div class="chat" v-if="isJoined">
      <form @submit.prevent="createMessage">
        <div class="columns is-mobile">
          <div class="column is-four-fifths"><input placeholder="message" class="input" type="text" v-model="newMessageBody"></div>
          <div class="column is-one-fifth"><button class="button is-primary is-fullwidth" :disabled="!newMessageBody">send</button></div>
        </div>
      </form>

      <table class="table is-fullwidth">
        <tr v-for="message in messages" :key="message.id">
          <td>
            <div class="columns is-vcentered">
              <p class="column is-one-fifth">{{message.createdBy}}</p>
              <p class="column is-three-fifths">{{message.body}}</p>
              <p class="column has-text-right has-text-grey-light is-size-7">{{message.created_at | datetime}}</p>
            </div>
          </td>
        </tr>
      </table>
    </div>

    <div class="join" v-else>
      <h2 class="title is-2">Type your nickname and join!</h2>
      <form @submit.prevent="join">
        <div class="columns is-mobile">
          <div class="column is-four-fifths"><input placeholder="nickname" class="input" type="text" v-model="nickname"></div>
          <div class="column is-one-fifth"><button class="button is-primary is-fullwidth" :disabled="!nickname">join</button></div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import moment from 'moment'

export default {
  name: 'Room',
  filters: {
    datetime(s) {
      const m = new moment(s)
      return m.format('HH:mm:ss')
    },
  },
  data() {
    return {
      newMessageBody: '',
      nickname: '',
      isJoined: false,
      room: {
        id: 1,
        created_at: 'Mon Mar 02 2020 14:23:46 GMT+0900',
        name: 'hogehogehogehogehogehoge',
      },
      messages: [
        {
          id: 1,
          body: 'hello',
          createdBy: 'taro',
          created_at: 'Mon Mar 02 2020 14:23:46 GMT+0900',
        },
        {
          id: 2,
          body: 'hello',
          createdBy: 'jiro',
          created_at: 'Mon Mar 02 2020 14:23:46 GMT+0900',
        },
      ],
    }
  },
  props: {
    roomId: String,
  },
  methods: {
    createMessage() {
      console.log(this.newMessageBody)
      this.newMessageBody = ''
    },
    join() {
      console.log(this.nickname)
      this.isJoined = true
    },
  },
}
</script>
