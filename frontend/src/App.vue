<template>
  <div id="app">
    <login-window v-if="loginVis" @close="signIn" />

    <add-user
      v-if="addUserVis"
      @update="updateWindow"
      @close="closeAddUser"
      @resetHamburgerImage="resetHamburgerImage"
    />

    <administrator-window
      v-if="adminWindowVis"
      :login="login"
      @signOut="signOut"
      @switchModalWindow="switchModalWindow"
      @close="closeAdministratorWindow"
    />

    <div class="background-color">
      <div class="header">
        <svg class="header__logo">
          <text class="header__crds">
            <tspan fill="#000000">cloud</tspan>
            <tspan fill="#000000">suny</tspan>
          </text>
        </svg>

        <div
          class="header__wrapper"
          :class="{
            'header__wrapper_scroll-correction':
              loginVis || addUserVis || switchHamburgerImage,
          }"
        >
          <div
            v-if="isAuth"
            ref="hamburger"
            class="header__hamburger"
            :class="[
              switchHamburgerImage
                ? houseBackgroundClass
                : burgerBackgroundClass,
            ]"
            @click="switchMenu"
          ></div>

          <a
            v-if="!isAuth"
            class="header__button header__button_m"
            @click="toogleLoginWindow"
            >Login</a
          >

          <div v-else class="header__lgnBlock">
            <a
              class="header__button header__button_l"
              @click="toogleAddUserWindow"
              >+ Add user</a
            >

            <p class="header__text">{{ login }}</p>

            <a @click="signOut" class="header__button_s">
              <svg
                width="20"
                height="20"
                viewBox="0 0 20 20"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M12.9167 15H12.0834C11.8532 15 11.6667 15.1866 11.6667
                 15.4167V18.3334H1.66668V1.66668H11.6667V4.58336C11.6667 4.81348
                 11.8532 5.00004 12.0834 5.00004H12.9167C13.1468 5.00004 13.3334
                 4.81348 13.3334 4.58336V1.66668C13.3333 0.746211 12.5871 0 11.6667
                 0H1.66668C0.746211 0 0 0.746211 0 1.66668V18.3334C0 19.2538 0.746211
                 20 1.66668 20H11.6667C12.5871 20 13.3334 19.2538 13.3334 18.3333V15.4166C13.3333
                 15.1866 13.1468 15 12.9167 15Z"
                  fill="#001A34"
                />
                <path
                  d="M19.8649 9.69254L14.8649 5.10922C14.7429 4.99812 14.5659
                 4.96761 14.4157 5.03515C14.2643 5.10148 14.1667 5.25121 14.1667
                 5.4164V6.24976C14.1667 6.36777 14.2167 6.48047 14.3046 6.55941L17.2011
                 9.16644H5.41668C5.18637 9.16644 5 9.35281 5 9.58312V10.4164C5 10.6468
                 5.18637 10.8331 5.41668 10.8331H17.2011L14.3046 13.4402C14.2167 13.5191
                 14.1667 13.6318 14.1667 13.7498V14.5831C14.1667 14.7483 14.2643 14.8981
                 14.4157 14.9644C14.4694 14.9884 14.5268 14.9998 14.5834 14.9998C14.6855 14.9998
                 14.7864 14.9619 14.8649 14.8903L19.8649 10.307C19.9512 10.228 20 10.1166
                 20 9.99976C20 9.88297 19.9512 9.77148 19.8649 9.69254Z"
                  fill="#001A34"
                />
              </svg>
            </a>
          </div>
        </div>
      </div>
    </div>

    <div class="wrap-cards">
      <div class="display-card">
        <drow-cards
          v-for="card in cardMass"
          :key="card.id"
          :card="card"
          :badges="card.badges"
          :isAuth="isAuth"
          class="card"
        />
      </div>
    </div>

  </div>
</template>

<script>
import axios from 'axios';
import DrowCards from './components/DrawCards.vue';
import AddUser from './components/ModalInitCard.vue';
import LoginWindow from './components/LoginWindow.vue';
import AdministratorWindow from './components/AdministratorWindow.vue';
import { valuesName, footerHeight, scrollLockUp } from './config/config';

export default {
  name: 'App',
  components: {
    AddUser,
    AdministratorWindow,
    DrowCards,
    LoginWindow,
  },
  data() {
    return {
      cardMass: {},

      burgerBackgroundClass: 'header__hamburger_burger-background',
      houseBackgroundClass: 'header__hamburger_house-background',

      addUserVis: false,
      loginVis: false,
      adminWindowVis: false,

      login: '',
      isAuth: false,

      limit: 50,
      offset: 0,
      scrollOffset: 0,

      switchHamburgerImage: false,

      timerId: null,

      loadLock: false,
      scrollLock: false,

      stopTick: false,

      errStatus: '',
    };
  },
  computed: {
    scrollHeight() {
      return window.innerHeight + footerHeight;
    },

  },
  methods: {
    // modal win
    toogleLoginWindow() {
      this.loginVis = !this.loginVis;

      this.toogleScroll(this.loginVis);

      this.toogleOnTimerUpdate(this.loginVis);
    },

    toogleAddUserWindow() {
      this.addUserVis = !this.addUserVis;

      this.toogleScroll(this.addUserVis);

      this.toogleOnTimerUpdate(this.addUserVis);
    },
    closeAddUser() {
      this.addUserVis = false;
      this.stopTick = false;

      this.toogleScroll();

      this.toogleImageHamburger(this.addUserVis);

      this.onTimerUpdate();
    },

    // login and logout methods
    signIn(isAuth, login) {
      this.toogleScroll();

      this.isAuth = isAuth;
      this.login = login;

      this.loginVis = !this.loginVis;

      this.stopTick = false;

      this.onTimerUpdate();
    },
    signOut() {
      axios
        .post('/api/v1/auth/sign_out', {})
        .then(() => {
          localStorage.removeItem(valuesName.localStorageLoginValue);

          this.isAuth = false;
          this.login = '';
        })
        .catch(() => alert('problem with sign out'));
    },

    toogleScroll(hideScroll) {
      const height = hideScroll ? '100vh' : '';
      const visible = hideScroll ? 'hidden' : 'visible';

      document.body.style.height = height;
      document.body.style.overflowY = visible;
    },

    // update and load cards
    initMass(response) {
      if (response.data.data === undefined) {
        return;
      }

      response.data.data.forEach(
        (item) => {
          if (this.cardMass[item.id] === undefined) {
            this.$set(this.cardMass, item.id, item);
          }
        },
      );
    },

    onTimerUpdate() {
      if (this.stopTick === true) {
        return;
      }

      this.timerId = setTimeout(() => {
        this.updateWindow();

        this.onTimerUpdate();
      }, 10000);
    },
    toogleOnTimerUpdate(stopTimer) {
      if (!stopTimer) {
        this.onTimerUpdate();
        return;
      }

      this.stopTick = true;

      clearTimeout(this.timerId);
    },

    updateWindow() {
      if (this.loadLock === true) {
        return;
      }

      const useCheckLoadImg = false;
      this.axiosMethod(this.offset, useCheckLoadImg);
    },

    handleScroll() {
      if (document.body.scrollHeight - this.scrollHeight - window.scrollY >= scrollLockUp) {
        this.scrollLock = false;
      }

      if (this.loadLock === true || this.scrollLock === true) {
        return;
      }

      if (document.body.scrollHeight - this.scrollHeight - window.scrollY <= 0) {
        const useCheckLoad = true;

        this.axiosMethod(this.scrollOffset, useCheckLoad);

        this.scrollLock = true;
      }
    },

    // Universal method for interaction with server
    axiosMethod(offsetChoose, useCheckLoadImg = false) {
      this.loadLock = true;

      const func = useCheckLoadImg ? this.checkLoadImg : this.initMass;

      axios
        .post('/api/v1/user/list/with_badges', {
          limit: this.limit,
          offset: offsetChoose,
        })
        .then((response) => {
          this.checkLoginUser(response);

          func(response);
        })
        .catch((error) => {
          alert(error.response.status);
        })
        .then(() => {
          this.loadLock = false;
        });
    },

    checkLoginUser(response) {
      this.isAuth = (response.headers['x-admin'] === 'true');

      if (localStorage.getItem(valuesName.localStorageLoginKey) !== null) {
        this.login = localStorage.getItem(valuesName.localStorageLoginKey);
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

    resetHamburgerImage(resetImage) {
      this.switchHamburgerImage = resetImage;
    },

    toogleImageHamburger(isModalWin) {
      this.switchHamburgerImage = isModalWin;

      this.toogleScroll(isModalWin);
    },

    switchMenu() {
      switch (true) {
        case this.addUserVis:
          this.toogleImageHamburger(this.addUserVis = false);
          break;
        case this.adminWindowVis:
          this.toogleImageHamburger(this.adminWindowVis = false);
          break;
        default:
          this.toogleImageHamburger(this.adminWindowVis = true);
      }
    },

    closeAdministratorWindow() {
      this.toogleScroll(this.adminWindowVis = false);

      const isModalWindow = false;
      this.toogleImageHamburger(isModalWindow);
    },

    switchModalWindow() {
      this.adminWindowVis = false;

      this.toogleAddUserWindow();

      const isModalWindow = true;
      this.toogleImageHamburger(isModalWindow);
    },
  },
  beforeMount() {
    window.addEventListener('scroll', this.handleScroll);
  },
  mounted() {
    const useCheckLoad = true;
    this.axiosMethod(this.offset, useCheckLoad);

    this.onTimerUpdate();
  },
  destroyed() {
    window.removeEventListener('scroll', this.handleScroll);
  },
};
</script>

<style lang="scss" scoped>
@import url("https://fonts.googleapis.com/css2?family=Comfortaa:wght@300;400;500;600;700&display=swap");

#app {
  font-family: "Comfortaa", cursive;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;

  margin: 0 auto;

  min-width: 430px;
  user-select: none;
}

.header {
  display: flex;
  flex-direction: row;
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
    background-repeat: no-repeat;
    background-position: center;

    &_burger-background {
      background-image: url("/Burger.svg");
    }
    &_house-background {
      background-image: url("/MenuHouse.svg");
    }
  }
  &__lgnBlock {
    display: flex;
    align-items: center;
  }
  &__button {
    display: inline-block;
    border: 2px solid black;
    border-radius: 60px;

    background-color: #ffff;

    cursor: pointer;
    font-size: 22px;
    font-weight: 700;

    &_l {
      margin: 0 20px 0 0;

      width: 158px;
      height: 50px;

      line-height: 50px;

      &:hover {
        background-color: #b6b9bb;
      }
      &:active {
        background-color: #000000;
        color: #ffffff;
      }
    }
    &_m {
      margin: 0;

      width: 115px;
      height: 40px;

      color: #000000;
      line-height: 40px;
      &:hover {
        color: #4d7fff;
      }
    }
    &_s {
      cursor: pointer;
      &:hover path {
        fill: #4d7fff;
        stroke: #4d7fff;
      }
    }
  }
  &__text {
    margin: 0 7px 0 0;

    font-style: normal;
    font-weight: bold;
    font-size: 22px;
    line-height: 25px;
  }
  &_visible {
    display: none;
  }
  &__wrapper {
    position: relative;
    &_scroll-correction {
      margin: 0 15px 0 0;
    }
  }
}
.wrap-cards {
  padding: 108px 0 0 0;
}
.display-card {
  display: flex;
  flex-wrap: wrap;

  margin: auto;
  width: 1680px;

  .card {
    margin: 5px;
  }
}

.createBtn {
  width: 50px;
  height: 50px;

  background-color: #3498db;
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
.background-color {
  background: #ffff;
  box-shadow: 0px 2px 0px rgba(129, 129, 129, 0.25);
  position: fixed;
  z-index: 40;
  width: 100%;
}
@media (max-width: 1680px) {
  .display-card {
    width: 1260px;
  }
}
@media (max-width: 1260px) {
  .display-card {
    width: 840px;
  }
}
@media (max-width: 840px) {
  .display-card {
    width: 420px;
  }
}
@media (max-width: 650px) {
  .header {
    &__hamburger {
      display: block;
      position: relative;
    }
    &__lgnBlock {
      display: none;
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
