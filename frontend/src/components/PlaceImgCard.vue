<template>
  <div
    @mouseup="onMouseUp"
    @mousemove="onMouseMove"
    ref="placeImgCard"
    class="element-place"
  >

    <div
      v-for="item in items"
      :key="item.id"
      :id="item.id"
      class="element-place__items"
      :ref="`element`+item.id"
      @mousedown="onMouseDown"
      :style="'left:'+item.x+'px;'+' top:'+item.y+'px;'"
    >

      <img
        class="element-place__imgSize"
        :src="item.img"
      />

    </div>
  </div>
</template>

<script>

const mouseup = 'mouseup';

export default {
  props: {
    items: {
      type: Array,
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
    };
  },
  methods: {
    takePlaceCoords(coordX, coordY) {
      const imgZone = this.$refs.placeImgCard;

      const obj = {
        left: imgZone.offsetLeft,
        top: imgZone.offsetTop,
        right: imgZone.offsetWidth + imgZone.offsetLeft - coordX,
        bottom: imgZone.offsetHeight + imgZone.offsetTop - coordY,
      };

      return obj;
    },

    onMouseDown(event) {
      event.preventDefault();

      this.getCoords(event);
      this.limitsPlace = this.takePlaceCoords(this.shiftW, this.shiftH);

      this.isDrag = true;

      document.addEventListener(mouseup, this.onMouseUp);
    },
    onMouseUp() {
      this.isDrag = false;

      document.removeEventListener(mouseup, this.onMouseUp);
    },
    onMouseMove(event) {
      if (this.isDrag) this.move(event);
    },

    move(event) { // TODO: попробовать упростить
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
    },
    // при нажатии на элемент запоминаем координаты нажатия
    getCoords(event) {
      this.dragElement = event.target.parentElement;

      this.shiftX = event.clientX - this.dragElement.getBoundingClientRect().left;
      this.shiftY = event.clientY - this.dragElement.getBoundingClientRect().top;

      this.shiftW = this.dragElement.getBoundingClientRect().right - event.clientX;
      this.shiftH = this.dragElement.getBoundingClientRect().bottom - event.clientY;
    },
  },
  mounted() {
    this.$emit('initPlace');
  },
};
</script>

<style lang="scss" scoped>
.element-place {
  width: 400px;
  height: 500px;
  margin: 0;
  position: relative;

  &__items {
    position: absolute;
    border: solid 1px black;
    background-color: #cde3fc;
    box-sizing: border-box;
    width: 50px;
    height: 50px;
  }
  &__imgSize {
    width: 100%;
    height: 100%;
  }
}
</style>
