<template>
  <div class="background">
    <div class="background_close" @click="close"></div>
    <div class="modal-win">

      <header class="logo modal-win__header">
        <p class="logo__wrap-text">Create User</p>
      </header>

      <div class="modal-body">

        <input
          type="text"
          v-model="username"
          placeholder="Enter Login"
          class="modal-body__input"
          @click="isFileEmpty = false, isEmpty = false"
          maxlength="15"
        />

        <label for="file"
        class="modal-body__input label">
          Change file
        </label>
        <input
            type="file"
            id="file"
            ref="file"
            v-on:change="handleFileUpload"
            class="modal-body__visible"
            @click="isFileEmpty = false, isEmpty = false"
        />

        <button class="modal-body__button-style" @click="create">
          OK
        </button>

        <div v-if="isEmpty||isFileEmpty" class="invalidValue">
          <span class="invalidValue__content">Invalid username or password</span>
        </div>

      </div>

      <a class="modal-win__close-btn close-btn"
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
      username: '',
      file: '',
      imagePath: '',
      isEmpty: false,
      isFileEmpty: false,
      spaceReg: / /g,
    };
  },
  methods: {
    handleFileUpload() {
      // console.log(this.$refs.file);
      this.file = this.$refs.file;
      console.log(this.$refs.file.files);
    },
    create() {
      if (!this.checkInput(this.username)) { this.isEmpty = true; return; }
      if (!this.checkInput(this.file)) { this.isFileEmpty = true; return; }
      this.uploadImg();
    },
    checkInput(element) {
      let param = '';
      switch (true) {
        case (typeof (element) === 'string'):
          param = element.replace(this.spaceReg, '');
          break;
        case (typeof (element) === 'object'):
          param = element;
          break;
        default:
      }
      // this.name = this.username.replace(/ /g, ''); // TODO: добавить регулярку
      if (param === '') { return false; }
      return true;
      // this.uploadImg();
    },
    uploadImg() {
      const formData = new FormData();
      formData.append('data', this.file.files[0]);
      axios
        .post('/api/v1/image/upload', formData,
          {
            headers: {
              'Content-Type': 'multipart/from-data',
            },
          })
        .then((response) => {
          console.log(response);
          this.imagePath = response.data.data.name;
          console.log(this.imagePath);
          this.createUser();
        })
        .catch((error) => {
          console.log(error);
        });
    },
    createUser() {
      axios
        .post('/api/v1/user/create', {
          username: this.username,
          avatar: this.imagePath,
        })
        .then((response) => {
          // this.id = response.data.data.id;
          const newElem = response.data.data;
          console.log(newElem);
          this.$emit('update', newElem);
          this.close();
        })
        .catch((resp) => {
          this.info = resp;
        });
    },

    close() {
      this.$emit('close');
    },
    positFalse(chooseElem) {
      // chooseElem = !chooseElem;
      let el = chooseElem;
      el = false;
      console.log(el);
    },
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
  top: 71px;

  backdrop-filter: blur(15px);
  background-color: hsla(0, 0%, 51%, 0.452);

  &_close {
    width: 100%;
    height: 100%;
  }
}

.modal-win {

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

.modal-body {
  display: flex;
  flex-direction: column;
  align-items: center;

  margin: 20px 0;

  &__input {

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
  &__visible {
    display: none;
  }
  &__button-style {
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
}

.close-btn {

  width: 30px;
  height: 30px;

  background-color: rgb(246, 246, 246);
  border-radius: 50%;

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

.label {
  appearance: button;

  &:hover {
    background-color: #B6B9BB;
    }
  &:active {
    background-color: #000000;
    color: #FFFFFF;
    }
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
