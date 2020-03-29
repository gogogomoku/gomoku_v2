import Vuex from "vuex";
import Vue from "vue";
import { cloneDeep } from "lodash";
import axios from "axios";

Vue.use(Vuex);

const initialState = {
  message: null,
  errorResponse: "",
  httpEndpoint: process.env.VUE_APP_SERVER_HTTP || "http://localhost:4242",
  match: {
    matchId: -1,
    currentPlayerId: -1,
    suggestionTimer: 0,
    suggestorOn: undefined,
    suggestion: {
      x: -1,
      y: -1
    },
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
};

export default new Vuex.Store({
  strict: true,
  state: cloneDeep(initialState),
  getters: {},
  mutations: {
    // update state
    setError(state, payload) {
      state.errorResponse = payload;
    },
    getHome(state, payload) {
      state.message = payload;
    },
    getMatch(state, payload) {
      const {
        id,
        board: { tab },
        p1,
        p2,
        suggestion: { X, Y },
        history
      } = payload;
      const lastPlayer =
        !history || history.length === 0
          ? null
          : history.map(h => h.player.id)[history.length - 1];
      state.errorResponse = "";
      state.match.matchId = id;
      state.match.currentPlayerId = lastPlayer ? lastPlayer ^ 0x3 : 1;
      state.match.board.tab = cloneDeep(tab);
      state.match.size = tab.length;
      state.match.players.p1.isAi = p1.isAi;
      state.match.players.p2.isAi = p2.isAi;
      state.match.players.p1.captured = p1.captured;
      state.match.players.p2.captured = p2.captured;
      state.match.suggestion.x = X;
      state.match.suggestion.y = Y;
      state.match.history = cloneDeep(history);
    },
    newMatch(state, payload) {
      const {
        id,
        board: { tab },
        p1,
        p2,
        suggestion: { X, Y },
        winner
      } = payload;
      state.errorResponse = "";
      state.match.matchId = id;
      state.match.currentPlayerId = 1;
      state.match.board.tab = cloneDeep(tab);
      state.match.size = tab.length;
      state.match.players.p1.isAi = p1.isAi;
      state.match.players.p2.isAi = p2.isAi;
      state.match.suggestion.x = X;
      state.match.suggestion.y = Y;
      state.match.winner = winner;
    },
    makeMove(state, payload) {
      state.match.currentPlayerId = state.match.currentPlayerId ^ 0x3;
      const {
        board: { tab },
        history,
        p1: { captured: p1Captured },
        p2: { captured: p2Captured },
        suggestion: { X, Y },
        winner
      } = payload;
      state.errorResponse = "";
      state.match.board.tab = cloneDeep(tab);
      state.match.history = cloneDeep(history);
      state.match.players.p1.captured = p1Captured;
      state.match.players.p2.captured = p2Captured;
      state.match.suggestion.x = X;
      state.match.suggestion.y = Y;
      state.match.winner = winner;
    },
    undoLastMove(state, payload) {
      state.match.currentPlayerId = state.match.currentPlayerId ^ 0x3;
      const {
        board: { tab },
        history,
        p1: { captured: p1Captured },
        p2: { captured: p2Captured },
        suggestion: { X, Y },
        winner
      } = payload;
      state.errorResponse = "";
      state.match.board.tab = cloneDeep(tab);
      state.match.history = cloneDeep(history);
      state.match.players.p1.captured = p1Captured;
      state.match.players.p2.captured = p2Captured;
      state.match.suggestion.x = X;
      state.match.suggestion.y = Y;
      state.match.winner = winner;
    },
    clearMatch(state) {
      Object.assign(state, cloneDeep(initialState));
      state.message = { Message: "You have navigated back to Gomoku_v2 home!" };
    }
  },
  actions: {
    getMatch({ state, commit }, { matchId }) {
      axios
        .get(`${state.httpEndpoint}/match/${matchId}`)
        .then(response => {
          const { data, status } = response;
          if (status !== 200) {
            throw new Error(`Error ${status}`);
          }
          commit("getMatch", data);
        })
        .catch(err => {
          console.info("Error in getMatch(): ", err);
          commit("setError", `Could not get match ${parseInt(matchId)}`);
        });
    },
    clearMatch({ commit }) {
      commit("clearMatch");
    },
    getHome({ state, commit }) {
      return fetch(`${state.httpEndpoint}`)
        .then(response => response.json())
        .catch(err => {
          throw new Error(err);
        })
        .then(data => {
          // TODO: const { Message: message } = data.message;
          commit("getHome", data);
        })
        .catch(err => {
          console.info(`Error in getHome(): `, err);
          commit("setError", `Could not connect to server.`);
        });
    },
    newMatch({ state, commit }, { p1ai, p2ai }) {
      // TODO: Error handling
      let url = new URL(`${state.httpEndpoint}/match/new`);
      const params = { p1ai, p2ai };
      Object.keys(params).forEach(key =>
        url.searchParams.append(key, params[key])
      );
      return fetch(url)
        .then(response => response.json())
        .catch(err => {
          throw new Error(err);
        })
        .then(data => {
          commit("newMatch", data);
        })
        .catch(err => {
          console.info("Error in newMatch(): ", err);
          commit("setError", "Could not create a new match.");
        });
    },
    makeMove({ state, commit }, { posX, posY }) {
      console.info(`Making move at ${[posX, posY]}`);
      const url = new URL(
        `${state.httpEndpoint}/match/${state.match.matchId}/move`
      );
      const options = {
        method: "POST",
        cache: "no-cache",
        body: JSON.stringify({
          id: state.match.matchId,
          playerId: state.match.currentPlayerId,
          posX,
          posY
        })
      };
      return fetch(url, options)
        .then(response => response.json())
        .then(data => {
          commit("makeMove", data);
        })
        .catch(() => {
          commit(
            "setError",
            `Could not make move for player ${state.match.currentPlayerId} at position x: ${posX}, y: ${posY}.`
          );
        });
    },
    async undoMove({ state, commit }) {
      console.info(`Undoing move, if there is one.`);
      if (state.match.history.length) {
        commit(
          "undoLastMove",
          await new Promise(resolve => {
            fetch(`${state.httpEndpoint}/match/${state.match.matchId}/undo`, {
              method: "POST",
              cache: "no-cache"
            })
              .then(results => resolve(results.json()))
              .catch(() => {
                const { history } = state.match;
                const lastItem = history[history.length - 1];
                const { player, position } = lastItem;
                commit(
                  "setError",
                  `Could not undo move for player ${player.id} at position x: ${position.X}, y: ${position.Y}.`
                );
              });
          })
        );
      }
    }
  }
});
