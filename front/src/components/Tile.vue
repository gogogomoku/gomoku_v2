<template>
  <div>
    <div v-if="posY == 0">{{ posX }}</div>
    <div
      class="btn inline-block p-2"
      :class="{'suggestion': isSuggestion, 'bg-red': value === 1, 'bg-somegray': value === 2}"
      @click="makeMove({ posX, posY })"
    >{{ parseInt(value) }}</div>
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
  computed: {
    ...mapState(["match"])
  },
  methods: { ...mapActions(["makeMove"]) }
};
</script>

<style scoped>
/* Ugly, temporary */
.suggestion {
  @apply bg-green;
}
.suggestion:hover {
  @apply bg-lightgreen;
}
</style>
