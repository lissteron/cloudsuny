<template>
  <div class="background" @keypress.enter="create">
    <div class="background_close" @click="close('close')"></div>
    <div :class="{'modal-win_new-width':isNextStep}" class="modal-win">

      <header class="logo modal-win__header">
         <p class="logo__wrap-text">{{mobileText}}</p>
      </header>

      <div class="modal-body">

        <div
          v-if="isNextStep"
          ref="imgPlace"
          class="imgPlace">
        </div>

        <input
          type="text"
          v-model="username"
          placeholder="Enter Login"
          class="modal-body__input"
          @click="dropError"
          maxlength="15"
        />

        <label
          v-if="!isNextStep"
          for = "file"
          class="modal-body__input label">
          select photo
        </label>

        <input
          type="file"
          id="file"
          ref="file"
          v-on:change="handleFileUpload"
          class="modal-body__visible"
          @click="dropError"
        />

        <button v-if="isNextStep" class="modal-body__button-style" @click="create">
          add user
        </button>

        <div v-if="isEmpty||isFileEmpty" class="invalidValue">
          <span class="invalidValue__content">{{errorMsg}}</span>
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
import {
  widthImg, heightImg, spaceReg, stringArr, mobileSize,
} from '../config/config';

export default {
  computed: {
    mobileText() {
      return (this.width <= mobileSize) ? 'Add User' : 'Create User';
    },
  },
  data() {
    return {
      width: 0,
      username: '',
      file: '',
      imagePath: '',
      imageName: '',
      isEmpty: false,
      isFileEmpty: false,
      isNextStep: false,
      errorMsg: '',
      imgSize: null,
    };
  },
  methods: {
    // load img method
    handleFileUpload() {
      this.file = this.$refs.file;
      this.imageName = this.file.files[0].name;
      this.isNextStep = true;

      const img = new Image();
      img.onload = () => {
        this.getImgSize(img);
      };
      img.src = URL.createObjectURL(this.file.files[0]);
    },
    getImgSize(img) {
      this.imgSize = img;

      const imger = document.createElement('img');

      imger.src = img.src;
      imger.style.width = '100%';
      imger.style.height = '100%';

      this.$refs.imgPlace.appendChild(imger);
    },
    uploadImg() {
      if (this.imagePath !== '') { this.createUser(); return; }

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
          switch (true) {
            case response.data.data === null && response.data.errors[0].code === '0':
              this.errorMsg = stringArr.imgLoadErr;
              this.isFileEmpty = true;
              break;
            default:
              this.imagePath = response.data.data.name;
              this.createUser();
          }
        })
        .catch(() => {
          window.location.reload();
        });
    },

    create() {
      if (!this.checkInput(this.username)) {
        this.isEmpty = true;
        this.errorMsg = stringArr.entUsrname;
        return;
      }
      if (!this.checkInput(this.file)) {
        this.isFileEmpty = true;
        this.errorMsg = stringArr.chooseImg;
        return;
      }
      if (parseInt(this.imgSize.height, 10) < heightImg
        || parseInt(this.imgSize.width, 10) < widthImg) {
        this.errorMsg = stringArr.sizeErr;
        this.isFileEmpty = true;
        return;
      }

      this.uploadImg();
    },
    checkInput(element) {
      let param = '';

      switch (true) {
        case (typeof (element) === 'string'):
          param = element.replace(spaceReg, '');
          break;
        case (typeof (element) === 'object'):
          param = element;
          break;
        default:
      }

      if (param === '') { return false; }
      return true;
    },
    createUser() {
      axios
        .post('/api/v1/user/create', {
          username: this.username,
          avatar: this.imagePath,
        })
        .then((response) => {
          const newElem = response.data.data;

          this.imagePath = '';
          this.isFileEmpty = true;

          this.$emit(stringArr.update, newElem);
          this.close(stringArr.ok);
        })
        .catch((resp) => {
          const desc = JSON.parse(resp.response.request.response);
          this.errorMsg = desc.details[0].description;
          this.isEmpty = true;
        });
    },

    close(btnName) {
      this.$emit(stringArr.cls, btnName);
    },

    updateWidth() {
      this.width = window.innerWidth;
    },
    dropError() {
      this.isFileEmpty = false;
      this.isEmpty = false;
      this.errorMsg = '';
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

.modal-win {

  position: absolute;
  background-color: #fafafa;
  width: 402px;
  height: 421px;

  z-index: 30;
  border-radius: 5%;
  box-shadow: 0 30px 60px 0 rgba(0, 0, 0, 0.5);
  &_new-width {
    height: 550px;
  }

  &__header {
    display: flex;
    justify-content: start;

    margin: 20px 0 0 0;
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

  margin: 35px 0;

  &__input {

    font-family: 'Roboto', sans-serif;
    display: inline-block;

    background: #FFFFFF;

    box-sizing: border-box;
    border: 1px solid #000000;
    -webkit-border-radius: 4px;
    border-radius: 4px;
    outline: none;

    margin: 0 0 15px 0;
    padding: 7px 17px;
    width: 350px;
    height: 40px;

    background-color: #f6f6f6;
    color: #0d0d0d;

    text-align: left;
    text-decoration: none;
    font-size: 15px;
  }
  &__text {
    visibility: hidden;
  }
  &__visible {
    display: none;
  }
  &__button-style {
    font-family: 'Roboto', sans-serif;
    display: inline-block;

    padding: 15px 80px;
    margin:  0 0 20px 0;
    width:  350px;
    height: 51px;
    outline: none;

    border: none;
    -webkit-border-radius: 6px;
    border-radius: 6px;

    background-color: #000000;
    color: white;
    font-size: 13px;
    font-weight: 900;
    text-transform: uppercase;
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
  background: #0870D1;
  border: 2px solid #0870D1;

  text-transform: uppercase;
  font-weight: 900;
  font-size: 20px;
  line-height: 23px;
  color: #FFFFFF;
  text-align: center;
  border-radius: 6px;

  &:hover {
    background-color: #B6B9BB;
    border-color:     #B6B9BB;
    }
  &:active {
    background-color: #000000;
    border-color:     #000000;
    color:            #FFFFFF;
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
.visible-button {
  display: none;
}
.imgPlace {
  width: 140px;
  height: 240px;
  margin: 0 0 20px 0;
}
@media(max-width: 650px){
  .background{
    align-items: flex-start;
    background-color: #fff;
    z-index: 50;
  }

  .modal-win {
    position: absolute;
    z-index: 60;
    width: 402px;
    height: 453px;
    background: #FFFFFF;
    border-radius: 20px;

    box-shadow: 0 0 0 0;

    &__header {
      margin: 40px 0 26px 0;
    }
    &__close-btn {
      display: none;
    }
  }
  .modal-body{
    &__button-style {
      font-style: normal;
      font-weight: 900;
      font-size: 24px;
      line-height: 24px;
      &:hover{
        background: #CFCFCF;
      }
    }
  }
  .visible-button {
  display: block;
  }
  .invalidValue {

    &__content {

      font-family: 'Roboto', sans-serif;
      font-style: normal;
      font-weight: 900;
      font-size: 24px;
      line-height: 24px;

    }
  }
}
</style>
