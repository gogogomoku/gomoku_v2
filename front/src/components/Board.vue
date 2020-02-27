<template>
  <div class="tab" v-if="match && match.board && match.board.tab && match.board.tab.length">
    <div class="row" v-for="(line, posY) in match.board.tab" :key="posY">
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

<style scoped>
.tab {
  width: 800px;
  padding: 10px;
  margin: 0 auto;
  flex-grow: 0;
  display: flex;
  flex-wrap: wrap;
  box-sizing: border-box;
}

.row {
  /* border: 1px solid red; */
  flex-basis: 100%;
  display: flex;
  align-items: stretch;
  justify-content: space-between;
}
</style>