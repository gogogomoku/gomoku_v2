<template>
  <div class="h-screen flex flex-wrap justify-center items-center text-center">
    <h1 class="flex-none min-w-full py-2 px-4 text-2xl text-white-whip">Welcome to Gomoku_v2!</h1>
    <pre v-if="message" class="flex-none min-w-full py-2 px-4 text-white-whip">{{ message.Message }}</pre>
    <button class="btn" @click="newMatch">New match</button>
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

