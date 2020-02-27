<template>
  <div class="hello">
    <h1>Gomoku_v2 Match {{match.id}}</h1>
    <div class="playerWrapper">
      <div class="playerBox">
        <p>p1</p>
        <p>AiStatus: {{ parseInt(match.players.p1.aiStatus) }}</p>
        <p>Captured stones: {{ parseInt(match.players.p1.captured) }}</p>
      </div>
      <div class="playerBox">
        <p>p2</p>
        <p>AiStatus: {{ parseInt(match.players.p2.aiStatus) }}</p>
        <p>Captured stones: {{ parseInt(match.players.p2.captured) }}</p>
      </div>
    </div>
    <Board v-if="match && match.board && match.board.tab && match.board.tab.length" />
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import Board from "./Board.vue";

export default {
  name: "Match",
  components: { Board },
  methods: {
    // local methods here
    ...mapActions(["newMatch"])
  },
  computed: {
    // local computed here
    ...mapState(["match"])
  }
};
</script>

<style scoped>
.playerWrapper {
  display: flex;
}
.playerClass {
  display: flex;
}
</style>
