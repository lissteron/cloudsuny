<template>
  <div id="app">
    <login v-if="lgvisible" @close="lgvisible=!lgvisible"/>
    <button @click="lgvisible=!lgvisible">Login</button>
    <button @click="deleteAll">del</button>
    <modal-win :adress="adress" @update='updateWin' @close="showModal" v-if="visible"/>
    <div class="display-card">
      <drow-cards
        v-for="card in cardMass"
        :key="card.id"
        :card="card"
        :adress="adress"
        class="card"
      />
    </div>

    <!-- <img alt="Vue logo" src="./assets/logo.png" />
    <create-user :adress="adress" @returnId="newValue" msg="Welcome to Your Vue.js App" />
    <create-bage :id="id" :adress="adress"/>
    <update-users :adress="adress"/> -->
    <!-- отображаем тут компонент, для которого совпадает маршрут -->
  </div>
</template>
<script>
import axios from 'axios';
import DrowCards from './components/DrawCards.vue';
import ModalWin from './components/ModalInitCard.vue';
import Login from './components/LoginWin.vue';
import obj from '../public/config/http';

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
      adress: '',
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
      // this.cardMass.push({ });
      this.visible = !this.visible;
    },
    deleteAll() {
      let elem = null;
      while (this.cardMass.length !== 0) {
        elem = this.cardMass.pop();
        console.log(elem.id);
        axios
          .post(`${this.adress}/api/v1/user/delete`, {
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
  // methods: {
  //   newValue(ider) {
  //     this.id = ider;
  //   },
  // },
  mounted() {
    this.adress = obj.adress;
    axios
      .post(`${this.adress}/api/v1/user/list/with_badges`, {})
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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 20px;
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
