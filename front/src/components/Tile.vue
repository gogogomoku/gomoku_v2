<template>
  <div>
    <div v-if="posY == 0">{{ posX }}</div>
    <div class="btn inline-block p-2" @click="makeMove({ posX, posY })">{{ parseInt(value) }}</div>
  </div>
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
