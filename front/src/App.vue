<template>
  <div id="app">
    <Arcade
      v-if="match.matchId === -1 && message && message.Message && message.Message.length"
      :message="message.Message"
    />

    <Match v-else-if="match.matchId > -1" />
  </div>
</template>

<script>
import Arcade from "./components/Arcade.vue";
import Match from "./components/Match.vue";
import { mapActions, mapState } from "vuex";

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
    ...mapState(["match", "message"])
  },
  methods: {
    ...mapActions(["getHome"])
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
