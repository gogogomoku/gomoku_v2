<template>
  <div class="tile" @click="makeMove({ posX, posY })">{{ parseInt(value) }}</div>
</template>

<script>
import { mapState, mapActions } from "vuex";
export default {
  name: "Tile",
  props: {
    value: Number,
    posX: Number,
    posY: Number
  },
  computed: mapState(["match"]),
  methods: { ...mapActions(["makeMove"]) }
};
</script>

<style>
.tile {
  /* border: 1px solid blue; */
  display: inline-block;
  margin: 0.25%;
  cursor: pointer;
}
</style>