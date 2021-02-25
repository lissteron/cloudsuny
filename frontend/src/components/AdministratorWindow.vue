<template>
  <div class="background">
    <div class="administrator-winindow">
      <p class="administrator-winindow__button">
        Signed in as <span class="text-style">{{ login }}</span>
      </p>

      <a class="administrator-winindow__button" @click="switchModalWindow"
        >Add user</a
      >

      <a @click="signOut" class="administrator-winindow__button"> Sign out </a>
    </div>
  </div>
</template>

<script>
import { mobileSize } from '../config/config';

export default {
  props: {
    login: {
      type: String,
      required: true,
    },
  },
  methods: {
    updateWidth() {
      if (window.innerWidth >= mobileSize) {
        this.closeAdministratorWindow();
      }
    },

    switchModalWindow() {
      this.$emit('switchModalWindow');
    },

    signOut() {
      this.$emit('signOut');

      this.closeAdministratorWindow();
    },

    closeAdministratorWindow() {
      this.$emit('close');
    },
  },
  mounted() {
    window.addEventListener('resize', this.updateWidth);
  },
  destroyed() {
    window.removeEventListener('resize', this.updateWidth);
  },
};
</script>

<style lang="scss" scoped>
.background {
  background-color: #fff;
  display: flex;
  justify-content: center;

  position: fixed;
  top: 70px;
  z-index: 20;

  width: 100%;
  height: 100%;
}

.administrator-winindow {
  display: flex;
  flex-direction: column;

  position: absolute;
  z-index: 30;

  width: 402px;
  height: 421px;

  &__button {
    cursor: default;
    margin: 0 auto;
    height: 45px;
    width: 381px;

    font-family: Comfortaa;
    font-style: normal;
    font-weight: 300;
    font-size: 18px;

    text-align: start;
    line-height: 50px;

    border-bottom: 1px solid rgba(129, 129, 129, 0.25);
    &:hover {
      color: #4d7fff;
    }
  }
}
.text-style {
  font-weight: bold;
}
</style>
