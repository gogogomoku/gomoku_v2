<template>
  <div
    class="btn rounded-center flex-none flex flex-wrap justify-center items-center resize-none m-0 p-0 h-10 w-10 border-solid border-1 border-grey-darkest content-center"
    :class="[isSuggestion ? 'bg-cyan-dark' : 'bg-gray-dark-4']"
    @click="sendMove({ posX, posY })"
    @mouseover="mouseOver"
    @mouseleave="mouseOut"
  >
    <div v-if="posY === 0" class="text-gray-600 -mt-12">{{ posX }}</div>
    <div class="h-4 w-3 p-4 rounded-full" :class="[stoneColor, stoneOpacity]" />
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  name: "Tile",
  props: {
    value: Number,
    posX: Number,
    posY: Number,
    isSuggestion: Boolean
  },
  data() {
    return {
      hovering: false
    };
  },
  computed: {
    stoneColor() {
      const {
        value,
        hovering,
        isSuggestion,
        match: { currentPlayerId: id }
      } = this;
      return !hovering && value === 0 && !isSuggestion
        ? ""
        : value === 1 || (value === 0 && (hovering || isSuggestion) && id === 1)
        ? "bg-black"
        : value === 2 || (value === 0 && (hovering || isSuggestion) && id === 2)
        ? "bg-white-whip"
        : "";
    },
    stoneOpacity() {
      const { value, hovering, isSuggestion } = this;
      return !hovering && value === 0 && !isSuggestion
        ? "opacity-0"
        : (!hovering && value > 0) || (hovering && isSuggestion)
        ? "opacity-100"
        : "opacity-50";
    },
    ...mapState(["match"])
  },
  methods: {
    mouseOver() {
      this.hovering = true;
    },
    mouseOut() {
      this.hovering = false;
    },
    async sendMove({ posX, posY }) {
      await this.makeMove({ posX, posY });
    },
    ...mapActions(["makeMove"])
  }
};
</script>

<style scoped>
/* .hasColumn {
  margin-top: 10px;
} */
</style>
