<template>
  <div class="background">

    <div class="login-window">

      <header class="logo login-window__header">
        <p class="logo__wrap-text">Sign In</p>
      </header>

      <div class="login-form login-window__main">
        <input
          class="login-form__inp-width"
          v-model="login"
          type="text"
          placeholder="login"
          required
        >
        <input
          class="login-form__inp-width"
          v-model="password"
          type="text"
          placeholder="password"
          required
        >

        <button
          class="button-style"
          name="form_auth_submit"
          @click="send"
        >SIGN IN</button>
      </div>

      <a
        class="login-window__close-btn close-btn"
        @click="close"
      >
        <img
          src="../assets/mark.svg"
          class="close-btn__img-size"
        />
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
    };
  },
  methods: {
    close() {
      this.$emit('close');
    },
    send() {
      console.log(this.adress);
      axios
        .post('/api/v1/auth/sign_in', {
          login: this.login,
          password: this.password,
        })
        .then((response) => { console.log(response); })
        .catch((error) => { console.log(`you have error == ${error}`); });
      this.close();
    },
  },
  mounted() {
    this.adress = process.env.SERVER_HOST;
  },
};
</script>

<style lang="scss" scoped>
@import url('https://fonts.googleapis.com/css2?family=Roboto:wght@300&display=swap');

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
}

.login-window {
  position: relative;
  background-color: #fafafa;
  width: 402px;
  height: 421px;

  border-radius: 10px;
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
    right: 5px;
    top: 5px;
  }
}
.logo {
  &__wrap-text {
    margin: 0 0 0 27px;
    font-size: 20px;

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

    margin: 0 0 20px 0;
    padding: 7px 17px;
    width: 350px;

    background-color: #f6f6f6;
    color: #0d0d0d;

    text-align: center;
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
  border: solid 1px #3aa1ca;

  padding: 6px;

  box-sizing: border-box;
  text-decoration: none;

  :active :hover {
    text-decoration: none;
  }

  &__img-size {
    width: 100%;
  }
}
.button-style {
  font-family: 'Roboto', sans-serif;
  display: inline-block;

  padding: 15px 80px;
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
</style>
