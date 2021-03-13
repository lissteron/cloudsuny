<template>
  <div class="background">
    <div class="background_close" @click="close"></div>
    <div class="login-window">
      <header class="logo login-window__header">
        <p class="logo__wrap-text">Sign In</p>
      </header>

      <form class="login-form login-window__main" @submit.prevent="send">
        <input
          class="login-form__input"
          v-model.trim="login"
          type="text"
          placeholder="login"
          autofocus
        />

        <input
          class="login-form__input"
          v-model.trim="password"
          type="password"
          placeholder="password"
        />

        <button class="button-style" type="submit">Sign in</button>

        <div v-if="errorMsg" class="invalid-value">
          <span class="invalid-value__content">{{ errorMsg }}</span>
        </div>
      </form>

      <a class="login-window__close-button close-button" @click="close">
        <svg class="close-button_color" xmlns="http://www.w3.org/2000/svg">
          <path d="M1 1L15 15M15 1L1 15" stroke="black" stroke-width="2" />
        </svg>
      </a>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { valuesName } from '../config/config';

export default {
  data() {
    return {
      login: '',
      password: '',
      isAuth: false,

      errorMsg: '',
    };
  },
  methods: {
    close() {
      this.$emit('close', this.isAuth, this.login);
    },

    send() {
      this.dropError();

      if (this.checkInput()) { return; }

      axios
        .post('/api/v1/auth/sign_in', {
          login: this.login,
          password: this.password,
        })
        .then(() => {
          localStorage.setItem(valuesName.localStorageLoginKey, this.login);
          this.isAuth = true;

          this.close();
        })
        .catch((error) => {
          this.errorMsg = error.response.statusText;
        });
    },

    checkInput() {
      switch (true) {
        case this.login === '':
          this.errorMsg = valuesName.errLogin;
          break;
        case this.password === '':
          this.errorMsg = valuesName.errPassword;
          break;
        default:
      }
      return this.errorMsg;
    },

    dropError() { this.errorMsg = ''; },
  },
};
</script>

<style lang="scss" scoped>
@import url("https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500;700;900&display=swap");

.background {
  display: flex;
  justify-content: space-around;
  align-items: center;

  position: fixed;
  z-index: 20;

  width: 100%;
  height: 100%;
  top: 70px;

  backdrop-filter: blur(15px);
  background-color: hsla(0, 0%, 51%, 0.452);

  &_close {
    width: 100%;
    height: 100%;
  }
}

.login-window {
  position: absolute;
  background-color: #fafafa;
  width: 402px;
  height: 421px;

  z-index: 30;
  border-radius: 5%;
  box-shadow: 0 30px 60px 0 rgba(0, 0, 0, 0.5);

  &__header {
    display: flex;
    justify-content: start;

    margin: 52px 0 36px 0;
  }
  &__main {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  &__close-button {
    position: absolute;
    right: 19px;
    top: 21px;
  }
}
.logo {
  &__wrap-text {
    margin: 0 0 0 27px;

    font-style: normal;
    font-weight: normal;
    font-size: 36px;
    line-height: 40px;

    cursor: default;
  }
}

.login-form {
  &__input {
    font-family: "Roboto", sans-serif;
    display: inline-block;

    background: #ffffff;

    box-sizing: border-box;
    border: 1px solid #000000;
    -webkit-border-radius: 4px;
    border-radius: 4px;
    outline: none;

    margin: 0 0 20px 0;
    padding: 7px 17px;
    width: 350px;
    height: 40px;

    background-color: #f6f6f6;
    color: #0d0d0d;

    text-align: left;
    text-decoration: none;
    font-size: 15px;
  }
}

// helper-classes
.close-button {
  width: 30px;
  height: 30px;

  background-color: rgb(246, 246, 246);
  border-radius: 50%;

  padding: 6px;

  box-sizing: border-box;
  text-decoration: none;
  &:hover path {
    stroke: #4d7fff;
  }
  &_color {
    width: 100%;
    height: 100%;
  }
}
.button-style {
  font-family: "Roboto", sans-serif;
  display: inline-block;

  padding: 15px 80px;
  margin: 0 0 20px 0;
  width: 350px;
  height: 51px;

  border: none;
  border-radius: 6px;
  outline: none;

  background: #000000;
  color: #ffffff;
  font-size: 13px;
  font-style: normal;
  font-weight: 900;
  text-decoration: none;
}
.invalid-value {
  display: flex;
  align-items: center;
  justify-content: center;

  width: 350px;
  height: 51px;

  border: 1px solid rgba(235, 0, 0, 0.85);
  box-sizing: border-box;
  border-radius: 6px;

  &__content {
    font-family: "Roboto", sans-serif;
    font-style: normal;
    font-weight: 900;
    font-size: 13px;
    color: rgba(235, 0, 0, 0.85);
    line-height: 15px;
    letter-spacing: -1px;
    text-transform: uppercase;
  }
}

@media (max-width: 650px) {
  .background {
    align-items: flex-start;
    background-color: #ffffff;
  }

  .login-window {
    height: 453px;
    background: #ffffff;
    border-radius: 20px;

    box-shadow: 0 0 0 0;

    &__header {
      margin: 60px 0 36px 0;
    }
    &__close-button {
      display: none;
    }
  }

  .button-style {
    font-size: 24px;
    line-height: 24px;

    &:hover {
      background: #cfcfcf;
    }
  }

  .invalid-value {
    &__content {
      font-size: 24px;
      line-height: 24px;
    }
  }
}
</style>
