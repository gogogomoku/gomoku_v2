<template>
  <div
    id="app"
    class="flex items-center justify-center w-screen h-screen antialiased bg-grey-darkest"
  >
    <Arcade
      v-if="match.matchId === -1 && message && message.Message && message.Message.length"
      :message="message.Message"
    />

    <Match v-else-if="match.matchId > -1" />
    <transition name="fade">
      <div
        v-if="errorResponse"
        class="fixed bottom-0 py-3 px-4 rounded-sm mb-8 shadow-lg bg-white border-solid border-gray-400 border-1 font-mono"
      >Error: Human-readable error message here</div>
    </transition>
  </div>
</template>

<script>
import Arcade from "./components/Arcade.vue";
import Match from "./components/Match.vue";
import { mapActions, mapState } from "vuex";
import "./assets/styles.css"; // Global styles implementing tailwind

export default {
  name: "App",
  components: {
    Arcade,
    Match
  },
  mounted() {
    this.getHome();
  },
  computed: {
    ...mapState(["match", "message", "errorResponse"])
  },
  methods: {
    ...mapActions(["getHome"])
  }
};
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.4s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
