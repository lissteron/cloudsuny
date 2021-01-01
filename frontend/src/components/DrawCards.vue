<template>
  <div>
    <div class="card-content">
      <picture>
        <img :src="adress+'/'+ card.avatar" class="card-content__img"/>
      </picture>
      <!-- передаем на аутсорс отрисовку солнышек и тучек -->
      <div class="wrapper">
        <PlaceImgCard :badges="badges" :adress="adress" class="wraper"/>
      </div>

      <footer class="footer">
        <p class="footer__text-wrap">{{ card.username }}</p>

        <button @click="addElemtoArr('sun')" class="footer__button">
          cloud
        </button>

        <button @click="addElemtoArr('cloud')" class="footer__button">
          sun
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
    adress: {
      type: String,
    },
  },
  data() { return { badges: [] }; },
  methods: {
    addElemtoArr(chooseImg) {
      axios
        .post(`${this.adress}/api/v1/badge/create`, {
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

  width: 400px;
  height: 550px;

  background-color: #0d56a6;

  &__img {
    width: 100%;
    height: 500px;
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
  background-color: #4186d3;

  &__text-wrap {
    width: 250px;
    text-align: left;
    &::after {
      display: block;
      content: "";
      height: 3px;
      width: 100%;
      background-color: #13518f;
    }
  }

  &__button {
    padding: 15px;
    outline: none;
    border: none;
    border-radius: 10px;
    background-color: #3498db;

    color: white;
    font-size: 16px;

    cursor: pointer;
    &:active {
      background-color: #274f7d;
    }
  }
}
</style>
