<template>
  <div id="app">
    <login v-if="lgvisible" @close="lgvisible=!lgvisible"/>
    <div class="header">
        <svg class="header__logo">
          <text x="0" y="50">
            <tspan fill="#55D1E7">cloud</tspan>
            <tspan fill="#FECE85">suny</tspan>
          </text>
        </svg>

        <a @click="lgvisible=!lgvisible" class="header__button">Login</a>
    </div>

    <!-- <button @click="deleteAll">del</button> -->
    <modal-win @update='updateWin' @close="showModal" v-if="visible"/>

    <div class="display-card">
      <drow-cards
        v-for="card in cardMass"
        :key="card.id"
        :card="card"
        class="card"
      />
    </div>
  </div>
</template>
<script>
import axios from 'axios';
import DrowCards from './components/DrawCards.vue';
import ModalWin from './components/ModalInitCard.vue';
import Login from './components/LoginWin.vue';

export default {
  name: 'App',
  components: {
    DrowCards,
    ModalWin,
    Login,
  },
  data() {
    return {
      cardMass: [],
      info: null,
      visible: false,
      lgvisible: false,
    };
  },
  methods: {
    updateWin(newElem) {
      this.cardMass.push(newElem);
    },
    showModal() {
      this.visible = !this.visible;
    },

    deleteAll() {
      let elem = null;
      while (this.cardMass.length !== 0) {
        elem = this.cardMass.pop();
        console.log(elem.id);
        axios
          .post('/api/v1/user/delete', {
            user_id: elem.id,
          })
          .then((response) => { console.log(response); })
          .catch((error) => { console.log(error); });
      }
    },
    deleteCard() {
      this.cardMass.pop();
    },

    initMass() {
      const bufArr = this.info.data.data;
      bufArr.forEach((item) => {
        this.cardMass.push(
          item,
        );
      });
    },
  },
  mounted() {
    axios
      .post('/api/v1/user/list/with_badges', {})
      .then((response) => {
        this.info = response;
        this.initMass();
      })
      .catch((error) => {
        this.info = error;
      });
  },
};
</script>

<style lang="scss" scoped>
@import url('https://fonts.googleapis.com/css2?family=Comfortaa:wght@300&display=swap');

#app {
  font-family: 'Comfortaa', cursive;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #F7F8FA;
  margin-top: 0;
}

.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;

  margin: 0px 0px 42px 0px;
  padding: 0px 70px;
  box-shadow: 0px 1px 0px rgba(129, 129, 129, 0.25);

  height: 70px;

  background-color: #FFFF;

  &__logo {
    font-size: 40px;
    font-weight: bold;
    cursor: default;

    height: inherit;
  }

  &__button {
    width: 115px;
    padding: 8px;

    border: 1px solid black;
    border-radius: 60px;

    background-color: #FFFF;

    cursor: pointer;
    font-size: 22px;
  }
}

.display-card {
  display: flex;
  flex-wrap: wrap;

  .card {
    margin: 5px;
  }
}

.createBtn{
  width: 50px;
  height: 50px;

  background-color:  #3498db;
  border: none;
  outline: none;
  border-radius: 50%;

  color: white;
  text-align: center;
  font-size: 20px;

  position: fixed;
  bottom: 10px;
  right: 10px;
  z-index: 10;
}
</style>
