<template>
  <div
    class="min-w-full min-h-full p-5"
    v-if="match && match.board && match.board.tab && match.board.tab.length"
  >
    <div
      class="row flex w-full justify-between align-stretch"
      v-for="(line, posY) in match.board.tab"
      :key="posY"
    >
      <Tile
        v-for="(tile, posX) in line"
        :key="posX + (match.board.tab.length * posY)"
        :value="match.board.tab[posY][posX]"
        :posX="posX"
        :posY="posY"
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
  methods: {},
  computed: {
    ...mapState(["match"])
  }
};
</script>
