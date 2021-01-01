<template>
  <div>
    <div class="card-content">
      <picture>
        <img :src="'/'+ card.avatar" class="card-content__img"/>
      </picture>
      <!-- передаем на аутсорс отрисовку солнышек и тучек -->
      <div class="wrapper">
        <PlaceImgCard :badges="badges" class="wraper"/>
      </div>

      <footer class="footer">
        <p class="footer__text-wrap">{{ card.username }}</p>

        <button @click="addElemtoArr('sun')" class="footer__button">
          <img src="sun.svg"/>
        </button>

        <button @click="addElemtoArr('cloud')" class="footer__button">
          <img src="cloud.svg"/>
        </button>

        <button @click="addElemtoArr('indian')" class="footer__button">
          <img src="indian.svg"/>
        </button>
      </footer>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import PlaceImgCard from './PlaceImgCardv2.vue';

export default {
  components: {
    PlaceImgCard,
  },
  props: {
    card: {
      type: Object,
      required: true,
    },
  },
  data() { return { badges: [] }; },
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
  },
  mounted() {
    console.log(this.card);
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

  width: 410px;
  height: 550px;

  background-color: #b3b3b3;

  &__img {
    width: 100%;
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
  justify-content: space-around;

  width: inherit;
  height: 50px;

  position: absolute;
  bottom: 0;
  background: none;

  &__text-wrap {
    width: 250px;
    text-align: left;
    padding: 0 0 0 10px;
  }

  &__button {
    outline: none;
    border: none;
    border-radius: 10px;
    background: none;

    cursor: pointer;
    &:active {
      background-color: #274f7d;
    }
  }
}
</style>
