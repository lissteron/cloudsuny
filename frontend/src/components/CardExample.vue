<template>
  <div>
    <transition name="modal">
      <ModalInitCard
        v-if="showModal"
        @set-ns-card="showModalVisible"
        @close ="close"
      />
    </transition>

    <div
      class="card-content"
      v-if="!showModal"
      v-bind:style="{ 'background-image': 'url(' + image + ')' }"
    >
      <!-- передаем на аутсорс отрисовку солнышек и тучек -->
      <PlaceImgCard
        ref="placeImgCard"
        :items="items"
        @initPlace="initPlace"
      />
      <footer class="footer">
        <p class="footer__text-wrap">{{ name }}</p>
        <button
          @click="addElemtoArr('./sun.png')"
          class="footer__button"
        >
          cloud
        </button>
        <button
          @click="addElemtoArr('./cloud.png')"
          class="footer__button"
        >
          sun
        </button>
      </footer>
    </div>
  </div>
</template>

<script>
import ModalInitCard from './ModalInitCard.vue';
import PlaceImgCard from './PlaceImgCard.vue';

export default {
  components: {
    ModalInitCard,
    PlaceImgCard,
  },
  data() {
    return {
      showModal: true,
      name: '',
      surname: '',
      idCount: 0,
      items: [],
      takeCoords: {},
      placeImg: {},
      image: './cloud.png',
    };
  },
  methods: {
    initPlace() {
      const imgZone = this.$refs.placeImgCard;

      this.takeCoords = {
        left: imgZone.offsetLeft,
        top: imgZone.offsetTop,
      };
      this.computedX = this.$refs.placeImgCard.left;
    },
    addElemtoArr(chooseImg) {
      this.idCount += 1;
      this.items.push({
        id: this.idCount,
        img: chooseImg,
        x: this.takeCoords.left,
        y: this.takeCoords.top,
      });
    },

    showModalVisible(name, surname) {
      this.showModal = !this.showModal;
      this.name = name;
      this.surname = surname;
    },
    close() {
      this.$emit('close');
    },
  },
};
</script>

<style lang="scss" scoped>
/* Обертка для содержимого */
.card-content {
  border: 1px solid black;
  width: 400px;

  background-size: contain;
  background-color: #0D56A6;
}

.footer {
  display: flex;
  justify-content: space-around;
  background-color: #4186D3;

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
      background-color: #274F7D;
    }
  }
}

.modal-enter-active {
  animation: modal-in 0.5s;
}
.modal-leave-active {
  animation: modal-in 0.5s reverse;
}
@keyframes modal-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
</style>
