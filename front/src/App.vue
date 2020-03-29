<template>
  <div id="app" class="flex flex-col align-center justify-center w-full antialiased">
    <transition name="fade">
      <div
        v-if="match.matchPending"
        class="fixed top-0 left-0 w-screen h-screen overflow-hidden bg-black opacity-75"
      >
        <div class="loader">Loading...</div>
      </div>
    </transition>
    <Match v-if="match.matchId > -1" />
    <Arcade v-else />
    <transition name="fade">
      <div
        v-if="errorResponse"
        class="fixed bottom-0 py-3 px-4 rounded-sm mb-8 shadow-lg bg-white border-solid border-gray-400 border-1 font-mono"
      >Error: {{ errorResponse }}</div>
    </transition>
    <div class="w-full flex -mt-12 py-3 px-4 text-white">
      <div class="flex">
        <img class="mr-2" src="./assets/github-logo.svg" width="15" height="15" />
      </div>
      <div>
        <a
          class="text-murky-yellow-gray no-underline align-baseline"
          href="https://github.com/gogogomoku/gomoku_v2"
        >@eskombro + @ekelen</a>
      </div>
    </div>
  </div>
</template>

<script>
import Arcade from "./pages/Arcade.vue";
import Match from "./pages/Match.vue";
import { mapActions, mapState } from "vuex";
import "./assets/styles.css"; // Global styles implementing tailwind
import "./assets/spinner.css";

export default {
  name: "App",
  components: {
    Arcade,
    Match
  },
  mounted() {
    const { id: routerId } = this.$route.params;
    const cb =
      routerId && routerId > -1
        ? this.getMatch.bind(null, { matchId: parseInt(routerId) })
        : this.getHome;
    cb();
  },
  computed: {
    ...mapState(["match", "errorResponse"])
  },
  methods: {
    updateMatch() {},
    ...mapActions(["getHome", "getMatch"])
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
