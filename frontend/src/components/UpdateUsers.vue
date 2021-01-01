<template>
  <div>
    <button @click="updateUser">Update</button>
    <div class="text">{{ info }}</div>
  </div>
</template>

<script>
import axios from 'axios';

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
    };
  },
  methods: {
    updateUser() {
      axios
        .post('/api/v1/user/list/with_badges', {})
        .then((response) => {
          this.info = response;
        })
        .catch((error) => {
          this.info = error;
        });
    },
  },
};
</script>

<style>
</style>
