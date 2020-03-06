/* eslint no-unused-vars: "off" */
import Vuex from "vuex";
import Vue from "vue";
import { cloneDeep } from "lodash";

Vue.use(Vuex);

export default new Vuex.Store({
  strict: true,
  state: {
    message: null,
    errorResponse: "",
    httpEndpoint: process.env.VUE_APP_SERVER_HTTP || "http://localhost:4242",
    match: {
      matchId: -1,
      currentPlayerId: -1,
      suggestionTimer: 0,
      suggestorOn: undefined,
      history: null,
      board: {
        size: 19,
        suggestedPosition: -1,
        tab: null,
        invalidMoves: [] // returned with prev move
      },
      players: {
        p1: {
          isAi: undefined,
          id: 1,
          captured: 0
        },
        p2: {
          isAi: undefined,
          id: 2,
          captured: 0
        }
      }
    }
  },
  getters: {},
  mutations: {
    // update state
    setError(state, payload) {
      state.errorResponse = payload;
    },
    getHome(state, payload) {
      state.message = payload;
    },
    newMatch(state, payload) {
      const {
        id,
        board: { tab },
        p1,
        p2
      } = payload;
      state.match.matchId = id;
      state.match.currentPlayerId = 1;
      state.match.board.tab = cloneDeep(tab);
      state.match.size = tab.length;
      state.match.players.p1.isAi = p1.isAi;
      state.match.players.p2.isAi = p2.isAi;
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
        await fetch(`${state.httpEndpoint}`).then(
          result => result.json(),
          async err => {
            commit("setError", err.toString());
            // await new Promise(r => setTimeout(r, 2000));
            // commit("setError", "");
          }
        )
      );
    },
    async newMatch({ state, commit }, { p1ai, p2ai }) {
      let url = new URL(state.httpEndpoint);
      const params = { p1ai, p2ai };
      Object.keys(params).forEach(key =>
        url.searchParams.append(key, params[key])
      );
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
