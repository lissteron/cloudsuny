<template>
  <div id="app">
    <login v-if="lgvisible" @close="signIn"/>
    <div class="header">
        <svg class="header__logo">
          <text x="0" y="50">
            <tspan fill="#000000">cloud</tspan>
            <tspan fill="#000000">suny</tspan>
          </text>
        </svg>
        <div class="buttons header__lgnBlock">
          <a @click="visible=!visible"
          class="header__button header__button_l"
          v-bind:class="{ 'header_visible': !isLogin }">
            + Add user
          </a>
          <a @click="lgvisible=!lgvisible"
          class="header__button header__button_m"
          v-bind:class="{'header_visible': isLogin}"
          >Login</a>
          <p class="header__text"
          v-bind:class="{header_visible:!isLogin}"
          >{{this.login}}</p>
          <a @click="signOut" class="header__button_s"
          v-bind:class="{ 'header_visible': !isLogin }">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M12.9167 15H12.0834C11.8532 15 11.6667 15.1866 11.6667
            15.4167V18.3334H1.66668V1.66668H11.6667V4.58336C11.6667 4.81348
            11.8532 5.00004 12.0834 5.00004H12.9167C13.1468 5.00004 13.3334
            4.81348 13.3334 4.58336V1.66668C13.3333 0.746211 12.5871 0 11.6667
            0H1.66668C0.746211 0 0 0.746211 0 1.66668V18.3334C0 19.2538 0.746211
            20 1.66668 20H11.6667C12.5871 20 13.3334 19.2538 13.3334 18.3333V15.4166C13.3333
            15.1866 13.1468 15 12.9167 15Z" fill="#001A34"/>
            <path d="M19.8649 9.69254L14.8649 5.10922C14.7429 4.99812 14.5659
            4.96761 14.4157 5.03515C14.2643 5.10148 14.1667 5.25121 14.1667
            5.4164V6.24976C14.1667 6.36777 14.2167 6.48047 14.3046 6.55941L17.2011
            9.16644H5.41668C5.18637 9.16644 5 9.35281 5 9.58312V10.4164C5 10.6468
            5.18637 10.8331 5.41668 10.8331H17.2011L14.3046 13.4402C14.2167 13.5191
            14.1667 13.6318 14.1667 13.7498V14.5831C14.1667 14.7483 14.2643 14.8981
            14.4157 14.9644C14.4694 14.9884 14.5268 14.9998 14.5834 14.9998C14.6855 14.9998
            14.7864 14.9619 14.8649 14.8903L19.8649 10.307C19.9512 10.228 20 10.1166 20 9.99976C20
            9.88297 19.9512 9.77148 19.8649 9.69254Z" fill="#001A34"/>
            </svg>

          </a>
        </div>
    </div>

    <!-- <button @click="deleteAll">del</button> -->
    <modal-win @update='updateWin' @close="showModal" v-if="visible"/>

    <div class="display-card">
      <drow-cards
        v-for="card in cardMass"
        :key="card.id"
        :card="card"
        :isLogin="isLogin"
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
      login: '',
      isLogin: false,
    };
  },
  methods: {
    updateWin(newElem) {
      this.cardMass.push(newElem);
    },
    showModal() {
      this.visible = !this.visible;
    },
    signIn(isLogin) {
      this.isLogin = isLogin;
      console.log(this.isLogin);
      this.lgvisible = !this.lgvisible;
    },

    signOut() {
      axios
        .post('/api/v1/auth/sign_out', {})
        .then((response) => {
          console.log(response);
          if (localStorage.removeItem('cloudsun') === undefined) {
            this.isLogin = false;
          }
        })
        .catch((error) => { console.log(`you have error == ${error}`); });
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
      if (this.info.headers['x-admin'] === 'true') {
        this.isLogin = true;
      }
      const bufArr = this.info.data.data;
      bufArr.forEach((item) => {
        this.cardMass.push(
          item,
        );
      });
    },
  },
  mounted() {
    if (localStorage.getItem('cloudsun') !== null) { this.login = localStorage.getItem('cloudsun'); }
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
@import url('https://fonts.googleapis.com/css2?family=Comfortaa:wght@300;400;500;600;700&display=swap');

#app {
  font-family: 'Comfortaa', cursive;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
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
  &__lgnBlock {
    display: flex;
    align-items: center;
  }
  &__button {

    display: inline-block;
    border: 2px solid black;
    border-radius: 60px;

    background-color: #FFFF;

    cursor: pointer;
    font-size: 22px;
    font-weight: 700;

    &_l{
      margin: 0 20px 0 0;

      width: 158px;
      height: 50px;

      line-height: 50px;

      &:hover {
        background-color: #B6B9BB;
      }
      &:active {
        background-color: #000000;
        color: #FFFFFF;
      }
    }
    &_m{
      margin: 0;

      width: 115px;
      height: 40px;

      color: #000000;
      line-height: 40px;
      &:hover {
        color: #4D7FFF;
      }
    }
    &_s {
      cursor: pointer;
      &:hover path{
        fill: #4D7FFF;
        stroke: #4D7FFF;
      }
    }
  }
  &__text {
    margin:0 7px 0 0;

    font-style: normal;
    font-weight: bold;
    font-size: 22px;
    line-height: 25px;
  }
  &_visible {
    display: none;
  }
}

.display-card {
  display: flex;
  flex-wrap: wrap;
  padding: 0 65px;

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
