<template>
  <div class="h-full flex flex-wrap p-10 text-white-whip">
    <h1 class="flex-none min-w-full">
      Gomoku_v2 Match {{match.matchId}} |
      <span
        @click="localClearMatch"
        class="text-cyan hover:cursor-pointer"
      >Home</span>
    </h1>
    <div class="flex-none min-w-full py-3 flex flex-wrap">
      <div class="flex flex-wrap">
        <p class="min-w-full">p1 (black stones)</p>
        <p class="min-w-full">Is AI: {{ `${match.players.p1.isAi}` }}</p>
        <p class="min-w-full">Captured stones: {{ parseInt(match.players.p1.captured) }}</p>
      </div>
      <div class="flex flex-wrap">
        <p class="min-w-full">p2 (white stones)</p>
        <p class="min-w-full">Is AI: {{ `${match.players.p2.isAi}` }}</p>
        <p class="min-w-full">Captured stones: {{ parseInt(match.players.p2.captured) }}</p>
      </div>
      <div class="flex flex-wrap">
        <button
          class="btn"
          @click="undoMove"
          :disabled="match.history === null || match.history.length === 0"
        >Undo</button>
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
    undoMove() {},
    localClearMatch() {
      this.clearMatch();
      this.$router.push("/");
    },
    ...mapActions(["newMatch", "undoMove", "clearMatch"])
  },
  computed: {
    // local computed here
    ...mapState(["match"])
  }
};
</script>
