<template>
  <div class="userBlock">
    <input type="text" placeholder="enter login" v-model="username" />
    <p>{{ getLogin }}</p>
    <button @click="createUser">Create User</button>
    <p>{{ info }}</p>
  </div>
</template>

<script>
import axios from 'axios';
// import { obj } from './config/http.js';

// console.log(obj);
export default {
  props: {
    adress: { // TODO: убрать, больше он не нужен.
      type: String,
      required: true,
    },
  },
  data() {
    return {
      info: null,
      point: {
        x: 0,
        y: 0,
      },
      type: 'string',
      avatar: 'string',
      username: 'string',
      id: '',
    };
  },
  methods: {
    createUser() {
      axios
        .post('/api/v1/user/create', {
          username: this.username,
          avatar: this.avatar,
        })
        .then((response) => {
          this.id = response.data.data.id;
          this.$emit('returnId', this.id);
        })
        .catch((resp) => {
          this.info = resp;
        });
    },
  },
  computed: {
    getLogin() {
      return this.username;
    },
  },
  mounted() {
  },

};
</script>

<style lang="scss" scoped>
.userBlock {
  width: 300px;
  height: 300px;
}
</style>
