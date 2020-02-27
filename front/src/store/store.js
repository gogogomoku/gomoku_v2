/* eslint no-unused-vars: "off" */
import Vuex from "vuex";
import Vue from "vue";
import { cloneDeep } from "lodash";

Vue.use(Vuex);

export default new Vuex.Store({
  strict: true,
  state: {
    message: null,
    httpEndpoint: process.env.VUE_APP_SERVER_HTTP || "http://localhost:4242",
    match: {
      matchId: -1,
      currentPlayerId: -1,
      suggestionTimer: 0,
      suggestorOn: false,
      history: null,
      board: {
        size: 19,
        suggestedPosition: -1,
        tab: null,
        invalidMoves: [] // returned with prev move
      },
      players: {
        p1: {
          aiStatus: 1,
          id: 1,
          captured: 0
        },
        p2: {
          aiStatus: 0,
          id: 2,
          captured: 0
        }
      }
    }
  },
  getters: {},
  mutations: {
    // update state
    getHome(state, payload) {
      state.message = payload;
    },
    newMatch(state, payload) {
      const {
        id,
        board: { tab }
      } = payload;
      state.match.matchId = id;
      state.match.currentPlayerId = 1;
      state.match.board.tab = cloneDeep(tab);
      state.match.size = tab.length;
    },
    makeMove(state, payload) {
      state.match.currentPlayerId = state.match.currentPlayerId ^ 0x3;
      const {
        board: { tab },
        history,
        p1: { captured: p1Captured },
        p2: { captured: p2Captured }
      } = payload;
      state.match.board.tab = cloneDeep(tab);
      state.match.history = cloneDeep(history);
      state.match.players.p1.captured = p1Captured;
      state.match.players.p2.captured = p2Captured;
    }
  },
  actions: {
    async getHome({ state, commit }) {
      commit(
        "getHome",
        await fetch(`${state.httpEndpoint}`).then(result => result.json())
      );
    },
    async newMatch({ state, commit }) {
      commit(
        "newMatch",
        await fetch(`${state.httpEndpoint}/match/new`).then(result =>
          result.json()
        )
      );
    },
    async makeMove({ state, commit }, { posX, posY }) {
      console.log(`Making move at ${[posX, posY]}`);
      commit(
        "makeMove",
        await fetch(`${state.httpEndpoint}/match/${state.match.matchId}/move`, {
          method: "POST",
          cache: "no-cache",
          body: JSON.stringify({
            id: state.match.matchId,
            playerId: state.match.currentPlayerId,
            posX,
            posY
          })
        }).then(results => results.json())
      );
    }
  }
});
