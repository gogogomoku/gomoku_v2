<template>
  <div class="hello">
    <h1>Gomoku_v2 Arcade</h1>
    <pre v-if="message">{{ message.Message }}</pre>
    <button
      v-if="message.Message && message.Message.length"
      @click="newMatch({p1ai: false, p2ai: false})"
    >New match</button>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  name: "Arcade",
  methods: {
    ...mapActions(["newMatch"])
  },
  computed: mapState(["message"])
};
</script>

<style scoped>
</style>
