<template>
  <div class="background">
    <div class="background_close" @click="close"></div>
    <div class="login-window">

      <header class="logo login-window__header">
        <p class="logo__wrap-text">Sign In</p>
      </header>

      <div class="login-form login-window__main">
        <input
          class="login-form__inp-width"
          v-model="login"
          type="text"
          placeholder="username"
          @click="isEmpty=true"
          required
        >
        <input
          class="login-form__inp-width"
          v-model="password"
          type="password"
          placeholder="password"
          @click="isEmpty=true"
          required
        >

        <button
          class="button-style"
          name="form_auth_submit"
          @click="send"
        >SIGN IN</button>

        <div v-if="!isEmpty" class="invalidValue">
          <span class="invalidValue__content">Invalid username or password</span>
        </div>
      </div>

      <a class="login-window__close-btn close-btn"
         @click="close">
        <svg class="close-btn_color" xmlns="http://www.w3.org/2000/svg">
        <path  d="M1 1L15 15M15 1L1 15" stroke="black" stroke-width="2"/>
        </svg>
      </a>

    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      login: '',
      password: '',
      isLogin: false,
      isEmpty: true,
    };
  },
  methods: {
    close() {
      console.log(this.isLogin);
      this.$emit('close', this.isLogin, this.login);
    },
    send() {
      console.log(`${this.password} and ${this.login}`);
      if (this.login === '' || this.password === '') { this.isEmpty = false; return; }
      console.log(this.adress);
      axios
        .post('/api/v1/auth/sign_in', {
          login: this.login,
          password: this.password,
        })
        .then((response) => {
          console.log(response); localStorage.setItem('cloudsun', this.login); this.isLogin = true;
          console.log(this.login);
          this.close();
        })
        .catch((error) => { console.log(`you have error == ${error}`); });
    },
  },
  mounted() {
    this.adress = process.env.SERVER_HOST;
  },
};
</script>

<style lang="scss" scoped>
@import url('https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500;700;900&display=swap');

.background {
  display: flex;
  justify-content: space-around;
  align-items: center;

  position: fixed;
  z-index: 20;

  width: 100%;
  height: 100%;
  top: 71px;

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
  &__close-btn {
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
  &__inp-width {

    font-family: 'Roboto', sans-serif;
    display: inline-block;

    background: #FFFFFF;

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
.close-btn {
  width: 30px;
  height: 30px;

  background-color: rgb(246, 246, 246);
  border-radius: 50%;
  // border: solid 1px #3aa1ca;

  padding: 6px;

  box-sizing: border-box;
  text-decoration: none;
  &:hover path{
    stroke: #4D7FFF;
  }
  &_color{
    width: 100%;
    height: 100%;
  }
}
.button-style {
  font-family: 'Roboto', sans-serif;
  display: inline-block;

  padding: 15px 80px;
  margin: 0 0 20px 0;
  width: 350px;
  height: 51px;

  border: none;
  -webkit-border-radius: 6px;
  border-radius: 6px;

  background-color: #000000;
  color: white;
  font-size: 13px;
  font-weight: 900;
  text-decoration: none;
}
.invalidValue {

  display: flex;
  align-items: center;
  justify-content: center;

  width: 350px;
  height: 51px;

  border: 1px solid rgba(235, 0, 0, 0.85);
  box-sizing: border-box;
  border-radius: 6px;

  &__content {

    font-family: 'Roboto', sans-serif;
    font-style: normal;
    font-weight: 900;
    font-size: 13px;
    color: rgba(235, 0, 0, 0.85);
    line-height: 15px;
    letter-spacing: -1px;
    text-transform: uppercase;
  }
}
</style>
