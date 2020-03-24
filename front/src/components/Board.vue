<template>
  <div
    class="min-w-full w-auto p-5"
    v-if="match && match.board && match.board.tab && match.board.tab.length"
  >
    <div
      class="row flex justify-center align-center p-0"
      v-for="(line, posY) in match.board.tab"
      :key="posY"
    >
      <div class="text-gray-600 w-5">{{ posY }}</div>
      <Tile
        v-for="(tile, posX) in line"
        :key="posX + (match.board.tab.length * posY)"
        :value="match.board.tab[posY][posX]"
        :posX="posX"
        :posY="posY"
        :isSuggestion="posY === match.suggestion.y && posX === match.suggestion.x"
      />
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import Tile from "./Tile.vue";

export default {
  name: "Board",
  components: { Tile },
  mounted() {
    const { matchId } = this.match;
    if (!this.$route.params.id)
      this.$router.push({ name: "match", params: { id: matchId } });
  },
  methods: {},
  computed: {
    ...mapState(["match"])
  }
};
</script>
