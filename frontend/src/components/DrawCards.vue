<template>
  <div>
    <div class="card-content">
      <picture>
        <img
          :src="'/' + card.avatar"
          class="card-content__img"
          @error="onError"
        />
      </picture>

      <div class="icons-wrap">
        <place-image-card :badges="badges" :isAuth="isAuth" />
      </div>

      <footer class="footer">
        <p class="footer__text-wrap">{{ card.username }}</p>

        <template v-if="isAuth">
          <button
            @click="badgeCreate('sun')"
            class="footer__button footer__button_sun"
          ></button>

          <button
            @click="badgeCreate('cloud')"
            class="footer__button footer__button_cloud"
          ></button>

          <button
            @click="badgeCreate('indian')"
            class="footer__button footer__button_indian"
          ></button>
        </template>
      </footer>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import PlaceImageCard from './PlaceImageCard.vue';
import { valuesName } from '../config/config';

export default {
  components: {
    PlaceImageCard,
  },
  props: {
    card: {
      type: Object,
      required: true,
    },
    isAuth: {
      type: Boolean,
      required: true,
    },
    badges: {
      type: Array,
      default() {
        return [];
      },
    },
  },
  data() {
    return {
      defaultImage: valuesName.defaultImage,
    };
  },
  methods: {
    badgeCreate(selectedImage) {
      axios
        .post('/api/v1/badge/create', {
          point: { x: 0, y: 0 },
          type: selectedImage,
          user_id: this.card.id,
        })
        .then((response) => {
          this.badges.push(response.data.data);
        })
        .catch((error) => {
          alert(error);
        });
    },

    onError(imageWithError) {
      const image = imageWithError;
      image.target.src = this.defaultImage;
    },
  },
};
</script>

<style lang="scss" scoped>
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

.icons-wrap {
  z-index: 10;

  position: absolute;
  top: 0;
  left: 0;
}

.footer {
  display: flex;
  align-items: center;

  width: inherit;
  height: 50px;

  position: absolute;
  bottom: 0;
  background: linear-gradient(
    0deg,
    rgba(0, 0, 0, 0.8) 0%,
    rgba(75, 75, 75, 0) 100%
  );

  &__text-wrap {
    width: 250px;

    font-style: normal;
    text-align: left;
    font-weight: bold;
    font-size: 20px;

    padding: 0 0 0 10px;

    color: #f0f8ff;
    cursor: default;
  }

  &__button {
    border: none;
    border-radius: 10px;
    outline: none;

    width: 50px;
    height: 50px;

    cursor: pointer;

    &_sun {
      background: url("/sun.svg");
      background-repeat: no-repeat;

      &:hover {
        background: url("/sunHover.svg");
        background-repeat: no-repeat;
      }
    }
    &_cloud {
      background: url("/cloud.svg");
      background-repeat: no-repeat;

      &:hover {
        background: url("/cloudHover.svg");
        background-repeat: no-repeat;
      }
    }
    &_indian {
      background: url("/indian.svg");
      background-repeat: no-repeat;

      &:hover {
        background: url("/indHover.svg");
        background-repeat: no-repeat;
      }
    }

    &:active {
      background-color: #274f7d;
    }
  }
}
</style>
