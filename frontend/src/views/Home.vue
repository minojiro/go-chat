<template>
  <div class="home">
    <h1 class="title is-1">chat</h1>
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
</template>

<script>
import moment from 'moment'

export default {
  name: 'Home',
  filters: {
    datetime(s) {
      const m = new moment(s)
      return m.format('HH:mm')
    },
  },
  data() {
    return {
      rooms: [
        {
          id: 1,
          created_at: 'Mon Mar 02 2020 14:23:46 GMT+0900',
          name: 'hogehogehogehogehogehoge',
        },
        {
          id: 2,
          created_at: 'Mon Mar 02 2020 14:23:46 GMT+0900',
          name: 'hogehogehogehogehogehoge',
        },
      ],
    }
  },
}
</script>
