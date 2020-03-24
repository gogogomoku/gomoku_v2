<template>
  <div class="h-full flex flex-wrap p-10 text-white-whip">
    <h1 class="flex-none w-full">Gomoku_v2 Match {{match.matchId}}</h1>
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
    ...mapActions(["newMatch", "undoMove"])
  },
  computed: {
    // local computed here
    ...mapState(["match"])
  }
};
</script>
