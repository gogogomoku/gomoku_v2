<template>
  <div class="h-full w-56-rem self-center flex flex-col p-10 text-white-whip">
    <h1 class="cursor-pointer" @click="localClearMatch">
      Gomoku_v2 Match {{match.matchId}} |
      <span class="text-cyan">Home</span>
    </h1>
    <div class="py-3 flex justify-end mb-4">
      <div class="flex flex-col mr-10">
        <p>p1 (black stones)</p>
        <p>Is AI: {{ `${match.players.p1.isAi}` }}</p>
        <p>Captured stones: {{ parseInt(match.players.p1.captured) }}</p>
      </div>
      <div class="flex flex-col mr-10">
        <p>p2 (white stones)</p>
        <p>Is AI: {{ `${match.players.p2.isAi}` }}</p>
        <p>Captured stones: {{ parseInt(match.players.p2.captured) }}</p>
      </div>
      <div class="flex flex-wrap">
        <button
          class="btn"
          @click="undoMove"
          :disabled="match.moveIsPending || match.history === null || match.history.length === 0"
        >Undo</button>
      </div>
    </div>
    <Board v-if="match && match.board && match.board.tab && match.board.tab.length" />
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import Board from "../components/Board.vue";

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
