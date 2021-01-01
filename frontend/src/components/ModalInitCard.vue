<template>
  <!-- <transition name="modal" :duration="{ enter: 800, leave: 800 }"> -->
  <div class="modal-mask">
    <div class="modal-wrapper">
      <div class="modal-header">
        <div name="header">Card Content</div>
      </div>

      <div class="modal-body">
        <div class="modal-body__label">Login</div>

        <input
          type="text"
          v-model="username"
          placeholder="Enter Login"
          class="modal-body__input"
          :class="{red: isEmpty}"
          @click="isEmpty = false"
          maxlength="15"
        />
        <label for="">
          File
          <input
            type="file"
            id="file"
            ref="file"
            v-on:change="handleFileUpload"
            @click="isLoad = false"
            v-bind:class="{red: isLoad}"
          />
        </label>
        <button class="modal-body__button button-style" @click="create">
          OK
        </button>
        <button class="modal-wrapper__button button-style" @click="close">
          X
        </button>
      </div>
    </div>
  </div>
  <!-- </transition> -->
</template>

<script>
import axios from 'axios';

export default {
  props: {
    adress: {
      type: String,
    },
  },
  data() {
    return {
      username: '',
      file: '',
      imagePath: '',
      isEmpty: false,
      isLoad: false,
      spaceReg: / /g,
    };
  },
  methods: {
    handleFileUpload() {
      this.file = this.$refs.file;
    },
    create() {
      if (!this.checkInput(this.username)) { this.isEmpty = true; return; }
      if (!this.checkInput(this.file)) { this.isLoad = true; }
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
      if (param === '') {
        alert('fuck off');
        return false;
      }
      return true;
      // this.uploadImg();
    },
    uploadImg() {
      if (this.imagePath !== '') {
        this.createUser();
        return;
      }
      const formData = new FormData();
      formData.append('data', this.file.files[0]);
      axios
        .post(`${this.adress}/api/v1/image/upload`, formData,
          {
            headers: {
              'Content-Type': 'multipart/from-data',
            },
          })
        .then((response) => {
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
        .post(`${this.adress}/api/v1/user/create`, {
          username: this.username,
          avatar: this.imagePath,
        })
        .then((response) => {
          // this.id = response.data.data.id;
          const newElem = response.data.data;
          console.log(newElem);
          this.$emit('update', newElem);
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
.modal-mask {
  position: fixed;
  z-index: 10;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(50, 74, 126, 0.267);
  //transition: all 2s;
}

.modal-wrapper {
  position: relative;
  width: 300px;
  margin: 10% auto;
  padding: 20px 30px;
  background-color: rgb(224, 224, 224);
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  //transition: all 2s;
  &__button {
    position: absolute;
    top: 0;
    right: 0;
    border-radius: 50%;
  }
}

.modal-body {
  display: flex;
  flex-direction: column;
  align-items: center;

  margin: 20px 0;

  &__input {
    width: 200px;
  }
  &__button {
    width: 50px;
    margin: 0 46px 0 0;
    align-self: flex-end;
  }
}

.button-style {
  background-color: #3498db;
  outline: none;
  border: none;
}

.red {
  border: 1px solid red;
}
</style>
