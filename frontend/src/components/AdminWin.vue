<template>
  <div class="bground" @keyup.enter="send">
    <div class="admin-win">
        <p class="admin-win__btn-text">
        Signed in as <span class="text-style">{{login}}</span>
        </p>

        <a class="admin-win__btn-text"
        @click="close('addUser')"
        >Add user</a>

        <a @click="close" class="admin-win__btn-text">
        Log out
        </a>
    </div>
  </div>
</template>

<script>
import { stringArr, mobileSize } from '../config/config';

export default {
  props: {
    login: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      isAddUser: false,
      width: 0,
    };
  },
  methods: {
    updateWidth() {
      this.width = window.innerWidth;
      if (this.width >= mobileSize) { this.close(stringArr.onlyClose); }
    },
    close(btnChoose) {
      switch (true) {
        case btnChoose === stringArr.addUser:
          this.$emit(stringArr.cls, stringArr.switchModal);
          break;
        case btnChoose === stringArr.onlyClose:
          this.$emit(stringArr.cls);
          break;
        default:
          this.$emit(stringArr.signOut);
          this.$emit(stringArr.cls);
      }
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
.bground {
    background-color: #fff;
    display: flex;
    justify-content: center;

    position: fixed;
    z-index:  20;

    width:  100%;
    height: 100%;
    top:    70px;

  &_close {
    width: 100%;
    height: 100%;
  }
}
.admin-win {
    display: flex;
    flex-direction: column;

    position: absolute;
    width: 402px;
    height: 421px;

    z-index: 30;

    &__btn-text{
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
          color: #4D7FFF;
        }
    }
}
.text-style{
    font-weight: bold;
}
</style>
