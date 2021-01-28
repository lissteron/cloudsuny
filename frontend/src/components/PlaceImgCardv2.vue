<template>
  <div
    @mouseup="onMouseUp"
    @mousemove="onMouseMove"
    ref="placeImgCard"
    class="element-place"
  >

    <div
      v-for="item in badges"
      :key="item.id"
      :id="item.id"
      class="element-place__items"
      @mousedown="onMouseDown"
      :style="'left:'+item.point.x+'px;'+' top:'+item.point.y+'px;'"
    >

      <img
        class="element-place__imgSize"
        :src="drawImg(item.type)"
      />

    </div>
  </div>
</template>

<script>
import axios from 'axios';

const mouseup = 'mouseup';

export default {
  props: {
    badges: {
      type: Array,
    },
    isLogin: {
      type: Boolean,
    },
  },
  data() {
    return {
      dragElement: null,
      isDrag: false,
      limitsPlace: {},
      shiftX: 0,
      shiftY: 0,
      shiftW: 0,
      shiftH: 0,
      getX: 0,
      getY: 0,
    };
  },
  methods: {
    takePlaceCoords(coordX, coordY) {
      const imgZone = this.$refs.placeImgCard;

      const obj = {
        left: imgZone.getBoundingClientRect().left,
        top: imgZone.getBoundingClientRect().top + window.pageYOffset,
        right: imgZone.offsetWidth + imgZone.getBoundingClientRect().left - coordX,
        bottom: imgZone.offsetHeight + imgZone.getBoundingClientRect().top
        + window.pageYOffset - coordY,
      };

      return obj;
    },

    onMouseDown(event) {
      event.preventDefault();
      if (this.isLogin === false) { return; }

      this.getCoords(event);
      this.limitsPlace = this.takePlaceCoords(this.shiftW, this.shiftH);

      this.isDrag = true;

      document.addEventListener(mouseup, this.onMouseUp);
    },
    onMouseUp() {
      if (this.dragElement !== null) {
        axios
          .post('/api/v1/badge/update', {
            id: this.dragElement.id,
            point: {
              x: this.getX,
              y: this.getY,
            },
          });
      }
      this.isDrag = false;

      document.removeEventListener(mouseup, this.onMouseUp);
    },
    onMouseMove(event) {
      if (this.isDrag) this.move(event);
    },

    move(event) {
      const newLocation = {
        x: this.limitsPlace.left,
        y: this.limitsPlace.top,
      };
      switch (true) {
        case event.pageX + this.shiftW > this.limitsPlace.right:
          newLocation.x = this.limitsPlace.right - this.shiftX - this.limitsPlace.left;
          break;
        case event.pageX - this.shiftX > this.limitsPlace.left:
          newLocation.x = event.pageX - this.shiftX - this.limitsPlace.left;
          break;
        default:
          newLocation.x = 0;
      }
      switch (true) {
        case event.pageY + this.shiftH > this.limitsPlace.bottom:
          newLocation.y = this.limitsPlace.bottom - this.shiftY - this.limitsPlace.top;
          break;
        case event.pageY - this.shiftY > this.limitsPlace.top:
          newLocation.y = event.pageY - this.shiftY - this.limitsPlace.top;
          break;
        default:
          newLocation.y = 0;
      }

      this.relocate(newLocation);
    },

    relocate(newLocation) {
      this.dragElement.style.left = `${newLocation.x}px`;
      this.dragElement.style.top = `${newLocation.y}px`;

      this.getX = newLocation.x;
      this.getY = newLocation.y;
    },
    // при нажатии на элемент запоминаем координаты нажатия
    getCoords(event) {
      this.dragElement = event.target.parentElement;

      this.getX = this.dragElement.getBoundingClientRect().left
      - this.$refs.placeImgCard.getBoundingClientRect().left;
      this.getY = this.dragElement.getBoundingClientRect().top
      - this.$refs.placeImgCard.getBoundingClientRect().top;

      this.shiftX = event.clientX - this.dragElement.getBoundingClientRect().left;
      this.shiftY = event.clientY - this.dragElement.getBoundingClientRect().top;

      this.shiftW = this.dragElement.getBoundingClientRect().right - event.clientX;
      this.shiftH = this.dragElement.getBoundingClientRect().bottom - event.clientY;
    },

    drawImg(imgType) {
      switch (imgType) {
        case 'sun':
          return './sunHov.svg';
        case 'cloud':
          return './cloudHvr.svg';
        default:
          return './indHvr.svg';
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.element-place {
  position: relative;
  width:  410px;
  height: 500px;
  margin: 0;

  &__items {
    position: absolute;
    box-sizing: border-box;
    width:  50px;
    height: 50px;
  }
  &__imgSize {
    width:  100%;
    height: 100%;
    background: none;
  }
}
</style>
