<template>
  <div class="background">

    <div class="login-window">

      <header class="logo login-window__header">
        <img
          src="../assets/cldsun.png"
          class="logo__wrap-img"
        />
        <p class="logo__wrap-text">SING IN</p>
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
        >Войти</button>
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
import obj from '../../public/config/http';

export default {
  data() {
    return {
      adress: '',
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
        .post(`${this.adress}/api/v1/auth/sign_in`, {
          login: this.login,
          password: this.password,
        })
        .then((response) => { console.log(response); })
        .catch((error) => { console.log(`you have error == ${error}`); });
      this.close();
    },
  },
  mounted() {
    this.adress = obj.adress;
  },
};
</script>

<style lang="scss" scoped>
.background {
  display: flex;
  justify-content: space-around;
  align-items: center;

  position: fixed;
  z-index: 20;

  width: 100%;
  height: 100%;

  background-color: #3aa1ca;
}

.login-window {
  position: relative;
  background-color: #fafafa;
  width: 450px;
  height: 400px;

  border-radius: 10px;
  box-shadow: 0 30px 60px 0 rgba(0, 0, 0, 0.3);

  &__header {
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 55px 0 30px 0;
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
  &__wrap-img {
    margin: 0 10px 0 0;
    height: 50px;
  }

  &__wrap-text {
    padding: 12px 0 6px 0;
    border-bottom: solid 5px #3aa1ca;

    font-size: 20px;
    cursor: default;
  }
}

.login-form {
  &__inp-width {
    display: inline-block;
    border: none;
    box-sizing: border-box;
    border-bottom: 2px solid #f6f6f6;

    -webkit-border-radius: 5px 5px 5px 5px;
    border-radius: 5px 5px 5px 5px;

    margin: 10px 0;
    padding: 15px 32px;
    width: 85%;

    background-color: #f6f6f6;
    color: #0d0d0d;

    text-align: center;

    text-decoration: none;
    font-size: 16px;

    -webkit-transition: all 0.3s ease-in-out;
    -moz-transition: all 0.3s ease-in-out;
    -ms-transition: all 0.3s ease-in-out;
    -o-transition: all 0.3s ease-in-out;
    transition: all 0.3s ease-in-out;
    &:focus {
      background-color: #fff;
      border-bottom: 2px solid #5fbae9;
    }
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
  color: black;

  :active :hover {
    text-decoration: none;
  }

  &__img-size {
    width: 100%;
  }
}
.button-style {
  display: inline-block;
  margin: 5px 20px 40px 20px;
  padding: 15px 80px;

  border: none;
  -webkit-border-radius: 5px 5px 5px 5px;
  border-radius: 5px 5px 5px 5px;

  background-color: #56baed;
  color: white;
  font-size: 20px;
  text-decoration: none;

  -webkit-box-shadow: 0 10px 30px 0 rgba(95, 186, 233, 0.4);
  box-shadow: 0 10px 30px 0 rgba(95, 186, 233, 0.4);
}
</style>
