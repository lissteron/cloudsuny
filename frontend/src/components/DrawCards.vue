<template>
  <div>
    <div class="card-content">
      <picture>
          <img :src="'/'+ card.avatar" class="card-content__img" @error="onError"/>
      </picture>
      <!-- передаем на аутсорс отрисовку солнышек и тучек -->
      <div class="wrapper">
          <PlaceImgCard :badges="badges" :isLogin="isLogin" class="wraper"/>
      </div>

      <footer class="footer"
      v-bind:class="{'footer-wrap': isLogin}">
          <p class="footer__text-wrap">{{ card.username }}</p>

          <button @click="addElemtoArr('sun')" class="footer__button"
          v-bind:class="{'footer__button-visible': !isLogin}">
            <img class= "" :src="sun" @mouseover="mouseOver('sun')"
            @mouseleave="mouseLeave('sun')" />
          </button>

          <button @click="addElemtoArr('cloud')" class="footer__button"
          v-bind:class="{'footer__button-visible': !isLogin}">
            <img class= "" :src="cloud" @mouseover="mouseOver('cloud')"
            @mouseleave="mouseLeave('cloud')"/>
          </button>

          <button @click="addElemtoArr('indian')" class="footer__button"
          v-bind:class="{'footer__button-visible': !isLogin}"
          >
            <img class= "" :src="indian" @mouseover="mouseOver('indian')"
            @mouseleave="mouseLeave('indian')"/>
          </button>
      </footer>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import PlaceImgCard from './PlaceImgCardv2.vue';
import imgErr from '../assets/cat1.jpg';

export default {
  components: {
    PlaceImgCard,
  },
  props: {
    card: {
      type: Object,
      required: true,
    },
    isLogin: {
      type: Boolean,
    },
  },
  data() {
    return {
      badges: [],
      sun: 'sun.svg',
      cloud: 'cloud.svg',
      indian: 'indian.svg',
    };
  },
  methods: {
    addElemtoArr(chooseImg) {
      axios
        .post('/api/v1/badge/create', {
          point: { x: 0, y: 0 },
          type: chooseImg,
          user_id: this.card.id,
        })
        .then((response) => { this.badges.push(response.data.data); });
    },

    mouseOver(el) {
      switch (el) {
        case 'sun':
          this.sun = 'sunHov.svg';
          break;
        case 'cloud':
          this.cloud = 'cloudHvr.svg';
          break;
        default:
          this.indian = 'indHvr.svg';
      }
    },
    mouseLeave(el) {
      switch (el) {
        case 'sun':
          this.sun = 'sun.svg';
          break;
        case 'cloud':
          this.cloud = 'cloud.svg';
          break;
        default:
          this.indian = 'indian.svg';
      }
    },

    onError(el) {
      const img = el;
      img.target.src = imgErr;
    },
  },
  mounted() {
    if (this.card.badges !== undefined) {
      this.badges = this.card.badges;
    }
  },
};
</script>

<style lang="scss" scoped>
/* Обертка для содержимого */
.card-content {
  position: relative;

  width:  410px;
  height: 550px;

  background-color: #b3b3b3;

  &__img {
    width:  100%;
    height: 550px;
  }
}

.wrapper {
  z-index: 10;

  position: absolute;
  top: 0;
  left: 0;
}

.footer {
  display: flex;

  width: inherit;
  height: 50px;

  position: absolute;
  bottom: 0;
  background: linear-gradient(0deg, rgba(0, 0, 0, 0.8) 0%, rgba(75, 75, 75, 0) 100%);

  &-wrap {
    justify-content: space-around;
    align-items: center;
  }
  &__text-wrap {
    width: 250px;

    font-style: normal;
    text-align: left;
    font-weight: bold;
    font-size: 20px;

    padding: 0 0 0 10px;

    color: aliceblue;
    cursor: default;
  }

  &__button {
    border: none;
    border-radius: 10px;
    background: none;
    outline: none;

    cursor: pointer;

    &-visible{ display: none;}

    &:active { background-color: #274f7d;}
  }
}
</style>
