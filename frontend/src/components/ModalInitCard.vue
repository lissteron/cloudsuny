<template>
  <div class="background">
    <div class="background_close" @click="close()"></div>
    <div :class="{ 'modal-window_new-width': secondStep }" class="modal-window">
      <header class="logo modal-window__header">
        <p class="logo__wrap-text">Add User</p>
      </header>

      <form
        class="modal-body"
        :class="{ 'modal-body_wrap': !secondStep }"
        @submit.prevent="create"
      >
        <div v-if="imagePath.src" ref="imgPlace" class="image-place">
          <img :src="imagePath.src" class="test" />
        </div>

        <input
          type="text"
          ref="inputText"
          v-model.trim="username"
          placeholder="Enter Username"
          class="modal-body__input"
          @click="dropError"
          maxlength="15"
        />

        <label v-if="!secondStep" for="file" class="modal-body__input label">
          select photo
        </label>

        <input
          type="file"
          id="file"
          ref="file"
          @change="handleFileUpload"
          @click="dropError"
          class="modal-body__visible"
        />

        <button
          class="modal-body__button_style"
          :class="{ 'modal-body__button_activate': secondStep }"
          type="submit"
          :disabled="disabled == true"
          :style="{ opacity: opacitySet }"
        >
          add user
        </button>

        <div v-if="errorMsg" class="invalid-value">
          <span class="invalid-value__content">{{ errorMsg }}</span>
        </div>
      </form>

      <a class="modal-window__close-button close-button" @click="close">
        <svg class="close-button_color" xmlns="http://www.w3.org/2000/svg">
          <path d="M1 1L15 15M15 1L1 15" stroke="black" stroke-width="2" />
        </svg>
      </a>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import {
  widthImg, heightImg, valuesName, mobileSize,
} from '../config/config';

export default {
  computed: {
    opacitySet() {
      return this.disabled ? '0.33' : '1';
    },
  },
  data() {
    return {
      username: '',

      emitForLargeScreen: true,

      file: '',
      avatar: '',

      imagePath: { src: '' },

      secondStep: false,
      errorMsg: '',
      disabled: true,
    };
  },
  methods: {
    // load img method
    handleFileUpload() {
      this.file = this.$refs.file;

      if (!this.identifyTypeImage(this.file.files[0].type)) { return; }

      this.secondStep = true;
      this.disabled = false;

      this.$refs.inputText.focus();

      const img = new Image();
      img.onload = () => { this.getImage(img); };
      img.src = URL.createObjectURL(this.file.files[0]);
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
          this.avatar = response.data.data.name;
          this.createUser();
        })
        .catch(() => {
          alert('some error with load Image');
        });
    },

    create() {
      if (this.identifyInput()) {
        return;
      }

      if (this.identifyImage()) {
        return;
      }

      if (!this.identifyTypeImage(this.file.files[0].type)) {
        return;
      }

      if (parseInt(this.imagePath.height, 10) < heightImg
        || parseInt(this.imagePath.width, 10) < widthImg) {
        this.errorMsg = valuesName.sizeErr;

        return;
      }

      this.uploadImg();
    },

    identifyInput() {
      this.errorMsg = (this.username === '') ? valuesName.entUsername : '';
      return this.errorMsg;
    },

    identifyImage() {
      this.errorMsg = (this.file === '') ? valuesName.chooseImg : '';
      return this.errorMsg;
    },

    identifyTypeImage(typeOfFile = '') {
      if (typeOfFile === undefined) {
        return false;
      }

      const str = typeOfFile.split('/');

      if (str.length <= 1) {
        return false;
      }

      if (str[0] !== valuesName.image) {
        this.errorMsg = valuesName.notImageLoad;

        return false;
      }

      if (str[1] !== valuesName.jpeg && str[1] !== valuesName.png) {
        this.errorMsg = valuesName.imgLoadErr;

        return false;
      }

      return true;
    },

    getImage(img) {
      this.imagePath = img;
    },

    createUser() {
      axios
        .post('/api/v1/user/create', {
          username: this.username,
          avatar: this.avatar,
        })
        .then(() => {
          this.$emit('update');

          this.close();
        })
        .catch((resp) => {
          const desc = JSON.parse(resp.response.request.response);

          if (!desc.details.length) {
            return;
          }

          this.errorMsg = desc.details[0].description;
        });
    },

    close() {
      this.$emit('close');
    },

    updateWidth() {
      switch (true) {
        case (window.innerWidth >= mobileSize && this.emitForLargeScreen):
          this.emitForLargeScreen = false;

          this.resetHamburgerImage(this.emitForLargeScreen);
          break;
        case (window.innerWidth <= mobileSize && !this.emitForLargeScreen):
          this.emitForLargeScreen = true;

          this.resetHamburgerImage(this.emitForLargeScreen);
          break;
        default:
      }
    },

    resetHamburgerImage(resetImage) {
      this.$emit('resetHamburgerImage', resetImage);
    },

    dropError() {
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
.test {
  width: 100%;
  height: 100%;
}
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

.modal-window {
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

.modal-body {
  display: flex;
  flex-direction: column;
  align-items: center;

  margin: 25px 0;

  &_wrap {
    margin: 70px 0;
  }

  &__input {
    font-family: "Roboto", sans-serif;
    display: inline-block;

    background: #ffffff;

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
  &__button_style {
    font-family: "Roboto", sans-serif;
    display: inline-block;

    padding: 15px 80px;
    margin: 0 0 20px 0;
    width: 350px;
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

.label {
  appearance: button;
  background: #0870d1;
  border: 2px solid #0870d1;

  text-transform: uppercase;
  font-weight: 900;
  font-size: 13px;
  line-height: 23px;
  color: #ffffff;
  text-align: center;
  border-radius: 6px;

  &:hover {
    background-color: #b6b9bb;
    border-color: #b6b9bb;
  }
  &:active {
    background-color: #000000;
    border-color: #000000;
    color: #ffffff;
  }
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

.image-place {
  width: 140px;
  height: 240px;
  margin: 0 0 20px 0;
}
@media (max-width: 650px) {
  .background {
    align-items: flex-start;
    background-color: #fff;
    z-index: 50;
  }

  .modal-window {
    position: absolute;
    z-index: 60;
    width: 402px;
    height: 453px;
    background: #ffffff;
    border-radius: 20px;

    box-shadow: 0 0 0 0;

    &__header {
      margin: 40px 0 26px 0;
    }
    &__close-button {
      display: none;
    }
  }
  .modal-body {
    &__button {
      &_style {
        font-style: normal;
        font-weight: 900;
        font-size: 24px;
        line-height: 24px;
      }
      &_activate {
        &:hover {
          background: #cfcfcf;
        }
      }
    }
  }

  .invalid-value {
    &__content {
      font-family: "Roboto", sans-serif;
      font-style: normal;
      font-weight: 900;
      font-size: 24px;
      line-height: 24px;
    }
  }
  .label {
    font-size: 20px;
  }
}
</style>
