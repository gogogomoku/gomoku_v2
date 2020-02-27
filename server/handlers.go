package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gogogomoku/gomoku_v2/arcade"
	"github.com/gogogomoku/gomoku_v2/board"
	"github.com/gogogomoku/gomoku_v2/player"
	"github.com/gorilla/mux"
)

type JsonMessage struct {
	Message string
}

type JsonMove struct {
	PlayerId int8 `json:"playerId"`
	PosX     int8 `json:"posX"`
	PosY     int8 `json:"posY"`
}

// GET /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("NEW HOME REQUEST")

	_ = json.NewEncoder(w).Encode(
		JsonMessage{Message: "Welcome to Gomoku... Use /new-match to create a match"},
	)
}

// GET /match/new
func NewMatchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("NEW NEW_MATCH REQUEST")

	new_match := *arcade.NewMatch()
	_ = json.NewEncoder(w).Encode(
		new_match,
	)
}

// GET /match/{id}
func GetMatchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("NEW GET_MATCH REQUEST")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	match := arcade.CurrentMatches.List[id]

	if match == nil {
		http.Error(w, "Bad request: match doesn't exist", http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(
		match,
	)
}

// POST /match/{id}/move
func PostMoveHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("NEW POST_MOVE REQUEST")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	match := arcade.CurrentMatches.List[id]

	if match == nil {
		http.Error(w, "Bad request: match doesn't exist", http.StatusBadRequest)
		return
	}
	params := JsonMove{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println("Error: ", err)
		http.Error(w, "Bad request: error in arguments", http.StatusBadRequest)
		return
	}
	log.Println(params)

	var player *player.Player
	if params.PlayerId == 1 {
		player = match.P1
	} else if params.PlayerId == 2 {
		player = match.P2
	} else {
		http.Error(w, "Bad request: bad player Id in arguments", http.StatusBadRequest)
		return
	}
	position := board.Position{X: params.PosX, Y: params.PosY}
	match.AddMove(player, &position)

	_ = json.NewEncoder(w).Encode(
		match,
	)

}
