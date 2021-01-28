<template>
  <div id="app">
    <login v-if="lgVis" @close="signIn"/>
    <add-user @update='updateCardsArr' @close="closeAddUser" v-if="addUserVis"/>
    <admin-win v-if="admwinVis" :login="login" @signOut="signOut" @close="showMenu"/>

    <div class="background-color">
      <div class="header">

        <svg class="header__logo">
          <text class="header__crds">
            <tspan fill="#000000">cloud</tspan>
            <tspan fill="#000000">suny</tspan>
          </text>
        </svg>

        <div class="header__wrapper">

          <div v-if="isLogin" ref="hamburger" class="header__hamburger" @click="showMenu">
             <img class="img" :src="require('../src/assets/'+imgHamName)" />
          </div>

          <a v-if="!isLogin"
            class="header__button header__button_m"
            @click="modalOpen('lgvisible')"
          >Login</a>

          <div v-if="isLogin" ref="lgnlock" class="header__lgnBlock">

            <a class="header__button header__button_l"
            @click="modalOpen('addUserVis')"
            >+ Add user</a>

            <p class="header__text">{{login}}</p>

            <a @click="signOut" class="header__button_s">

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
                 14.7864 14.9619 14.8649 14.8903L19.8649 10.307C19.9512 10.228 20 10.1166
                 20 9.99976C20 9.88297 19.9512 9.77148 19.8649 9.69254Z" fill="#001A34"/>
              </svg>

            </a>
          </div>

        </div>
      </div>
    </div>

    <div class="wrap-cards">
      <div class="display-card">
        <drow-cards
          v-for="card in pushCard"
          :key="card.id"
          :card="card"
          :isLogin="isLogin"
          class="card"
        />
      </div>
    </div>

    <img v-if ="errImg" :src="returnImgError">
  </div>
</template>

<script>
import axios from 'axios';
import DrowCards from './components/DrawCards.vue';
import AddUser from './components/ModalInitCard.vue';
import Login from './components/LoginWin.vue';
import AdminWin from './components/AdminWin.vue';
import { stringArr, footerHeight, scrollLockUp } from './config/config';

export default {
  name: 'App',
  components: {
    DrowCards,
    AddUser,
    AdminWin,
    Login,
  },
  computed: {
    scrollHeight() {
      return window.innerHeight + footerHeight;
    },
    returnImgError() {
      return `https://http.cat/${this.errStatus}`;
    },
  },
  data() {
    return {
      cardMass: {},
      pushCard: [],

      addUserVis: false,
      lgVis: false,
      admwinVis: false,

      login: '',
      isLogin: false,

      limit: 30,
      offset: 0,
      scrollOffset: 0,

      imgHamName: 'Burgerburger.svg',
      toogleimgHamName: true,

      timerUpdate: null,
      loadLock: false,
      scrollLock: false,
      stopTick: false,

      errImg: false,
      errStatus: '',
    };
  },
  methods: {
    switchImgHam() {
      return this.imgHamName;
    },
    // modal win
    modalOpen(element) {
      switch (element) {
        case stringArr.addUserVis:
          this.addUserVis = !this.addUserVis;
          break;
        default:
          this.lgVis = !this.lgVis;
      }

      this.clearIntervalUpdate();
      this.documentStyle('100vh', 'hidden');
    },
    closeAddUser() {
      this.addUserVis = !this.addUserVis;
      this.stopTick = false;

      this.documentStyle();
      this.toogleImgHam(this.addUserVis);
      this.setIntervalUpdate();
    },

    // login and logout methods
    signIn(isLogin, login) {
      this.isLogin = isLogin;
      if (login !== '') { this.login = login; }

      this.documentStyle();

      this.lgVis = !this.lgVis;

      this.stopTick = false;
      this.setIntervalUpdate();
    },
    signOut() {
      axios
        .post('/api/v1/auth/sign_out', {})
        .then(() => {
          if (localStorage.removeItem(stringArr.cookie) === undefined) {
            this.isLogin = false;
            this.login = '';
          }
        });
    },

    // block scroll and fix modalwin
    documentStyle(height = '', visible = '') {
      document.body.style.height = height;
      document.body.style.overflowY = visible;
    },

    // update and load cards
    initMass(response) {
      if (response.data.data === undefined) { return; }

      response.data.data.forEach(
        (item) => { this.cardMass[item.id] = item; },
      );

      this.pushCard.length = 0;

      Object.values(this.cardMass).forEach((item) => {
        this.pushCard.push(item);
      });
    },

    setIntervalUpdate() {
      if (this.stopTick === true) { return; }

      setTimeout(() => {
        this.updateWin();
        this.setIntervalUpdate();
      }, 10000);
    },
    clearIntervalUpdate() { this.stopTick = true; },

    updateWin() {
      if (this.loadLock === true) { return; }
      const useCheckLoadImg = false;
      const useCheckLogin = true;
      this.axiosMethod(this.offset, useCheckLoadImg, useCheckLogin);
    },
    updateCardsArr(newElem) {
      this.pushCard.push(newElem);
    },

    handleScroll() {
      if (document.body.scrollHeight - this.scrollHeight - window.scrollY >= scrollLockUp) {
        this.scrollLock = false;
      }
      if (this.loadLock === true || this.scrollLock === true) { return; }
      if (document.body.scrollHeight - this.scrollHeight - window.scrollY <= 0) {
        const useCheckLoad = true;

        this.axiosMethod(this.scrollOffset, useCheckLoad);

        this.scrollLock = true;
      }
    },

    // Universal method for interaction with serve
    axiosMethod(offsetChoose, useCheckLoadImg = false, useCheckLogin = false) {
      this.loadLock = true;
      let func = this.initMass;

      if (useCheckLoadImg) { func = this.checkLoadImg; }

      axios
        .post('/api/v1/user/list/with_badges', {
          limit: this.limit,
          offset: offsetChoose,
        })
        .then((response) => {
          if (useCheckLogin) { this.checkLoginUser(response); }
          func(response);

          this.loadLock = false;
        })
        .catch((error) => {
          this.loadLock = false;
          if (error.response.status !== undefined) {
            this.errStatus = error.response.status;
            this.toogleErrorImg(this.errStatus);
          }
        });
    },
    checkLoginUser(response) {
      this.isLogin = (response.headers['x-admin'] === 'true');

      if (localStorage.getItem(stringArr.cookie) !== null) {
        this.login = localStorage.getItem(stringArr.cookie);
      }
    },
    checkLoadImg(response) {
      this.offset = this.scrollOffset;

      let skipped = 0;
      if (response.data.skipped !== undefined) {
        skipped = parseInt(response.data.skipped, 10);
      }

      switch (true) {
        case (parseInt(response.data.count, 10) + skipped
            < parseInt(response.data.total, 10)):
          this.scrollOffset += response.data.count;
          break;
        default:
          this.scrollOffset = response.data.total;
      }

      this.initMass(response);
    },

    toogleErrorImg(errStatus) {
      if (this.pushCard.length === 0) {
        this.returnImgError(errStatus);
        this.errImg = true;
        return;
      }

      this.errImg = false;
    },
    toogleImgHam(isTrue) {
      switch (true) {
        case isTrue:
          this.imgHamName = 'MenuHouse.svg';
          break;
        default:
          this.imgHamName = 'Burgerburger.svg';
      }
    },

    showMenu(switchModal) {
      let imgChoose = true;
      switch (true) {
        case this.addUserVis === true:
          this.addUserVis = false;
          this.documentStyle();
          imgChoose = false;
          break;
        case switchModal === stringArr.switchModal:
          this.admwinVis = false;
          this.modalOpen(stringArr.addUserVis);
          break;
        case switchModal === stringArr.switchImg:
          this.admwinVis = false;
          imgChoose = false;
          break;
        case this.admwinVis:
          this.admwinVis = false;
          this.documentStyle();
          imgChoose = false;
          break;
        default:
          this.admwinVis = true;
          this.documentStyle('100vh', 'hidden');
      }

      this.toogleImgHam(imgChoose);
      this.switchImgHam();
    },
  },
  beforeMount() {
    window.addEventListener('scroll', this.handleScroll);
  },
  mounted() {
    const useCheckLoad = true;
    const useCheckLogin = true;
    this.axiosMethod(this.offset, useCheckLoad, useCheckLogin);

    this.setIntervalUpdate();
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

  margin: 0 auto;

  min-width: 430px;
}

.header {
  display: flex;
  flex-direction:  row;
  justify-content: space-between;
  align-items: center;

  margin: 0 auto;
  padding: 0px 25px;

  max-width: 1920px;
  height: 70px;

  &__logo {
    font-size: 40px;
    font-weight: bold;
    cursor: default;

    height: inherit;
  }
  &__crds {
      transform: translateY(50px);
  }
  &__hamburger {
    display: none;
    width: 40px;
    height: 30px;
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
  &__wrapper{
    position: relative;
  }
}
.wrap-cards {
  padding: 108px 0 0 0;
}
.display-card {
  display:   flex;
  flex-wrap: wrap;

  margin: auto;
  width:  1680px;

  .card {
    margin: 5px;
  }
}

.createBtn {
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
.disabled {
  pointer-events: none;
  opacity: 0.5;
  cursor: none;
}
.img {
  height: 100%;
  width: 100%;
}
.background-color {
  background: #FFFF;
  box-shadow: 0px 2px 0px rgba(129, 129, 129, 0.25);
  position: fixed;
  z-index: 40;
  width: 100%;
}
@media(max-width: 1680px){
  .display-card {
    width: 1260px;
  }
}
@media(max-width: 1260px){
  .display-card {
    width: 840px;
  }
}
@media(max-width: 840px){
  .display-card {
    width: 420px;
  }
}
@media(max-width: 650px){
  .header {

    &__hamburger {
      display: block;
      position: relative;
    }
    &__lgnBlock{
      display:none;
    }
    &__logo {
      font-size: 30px;
      font-weight: normal;
    }
    &__crds {
      transform: translateY(44px);
    }
    &__button {
       width: 88px;
       height: 30px;
       font-size: 18px;
       line-height: 30px;
     }
  }
}
</style>
