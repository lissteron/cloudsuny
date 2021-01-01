<template>
  <div class="createBlock">
    <p>
      {{ returnId }}
    </p>
    <input type="radio" name="type" @click="chooseValue('sun')" checked />sun
    <br />
    <input type="radio" name="type" @click="chooseValue('cloud')" />cloud <br />
    <input type="radio" name="type" @click="chooseValue('indian')" />indian
    <br />
    <button @click="createBage">CreateBage</button>
    <p>{{ typer }}</p>
    <p>{{ info }}</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  props: {
    id: {
      type: String,
      required: true,
    },
    adress: { // TODO: убрать, больше он не нужен.
      type: String,
      required: true,
    },
  },
  data() {
    return {
      info: '',
      point: {
        x: 0,
        y: 0,
      },
      typer: 'sun',
      bage_id: '',
    };
  },
  methods: {
    createBage() {
      axios
        .post('/api/v1/badge/create', {
          point: {
            x: 0,
            y: 0,
          },
          type: this.typer,
          user_id: this.id,
        })
        .then((response) => {
          this.info = response;
        })
        .catch((error) => { this.info = error; });
    },
    chooseValue(name) {
      this.typer = name;
    },
  },
  computed: {
    returnId() {
      return this.id;
    },
  },
};
</script>

<style>
.createBlock {
  width: 300px;
  height: 300px;
}
</style>
