package arcade

import (
	"errors"
	"fmt"
	"log"

	"github.com/gogogomoku/gomoku_v2/board"
	"github.com/gogogomoku/gomoku_v2/heuristic"
	"github.com/gogogomoku/gomoku_v2/player"
	pl "github.com/gogogomoku/gomoku_v2/player"
)

type Match struct {
	Id         int             `json:"id"`
	P1         *pl.Player      `json:"p1"`
	P2         *pl.Player      `json:"p2"`
	Suggestion *board.Position `json:"suggestion"`
	Winner     *pl.Player      `json:"winner"`
	Board      *board.Board    `json:"board"`
	History    []*board.Move   `json:"history"`
}

// A struct containing several simoultaneous matches
type Arcade struct {
	List    map[int]*Match
	Counter int
}

// Global object containing a reference to simoultaneous matches
var CurrentMatches = Arcade{
	List:    make(map[int]*Match),
	Counter: 0,
}

func GetOpponent(player *player.Player) *player.Player {
	if player == nil {
		return nil
	}
	match := CurrentMatches.List[player.MatchId]
	if player.Id == 1 {
		return match.P2
	} else {
		return match.P1
	}
}

// Creates a new match, stores it in Arcade map, returns it's address
func NewMatch(aiP1 bool, aiP2 bool) *Match {
	CurrentMatches.Counter++
	matchId := CurrentMatches.Counter
	p1 := pl.Player{Id: 1, OpponentId: 2, Captured: 0, IsAi: aiP1, MatchId: matchId}
	p2 := pl.Player{Id: 2, OpponentId: 1, Captured: 0, IsAi: aiP2, MatchId: matchId}
	match := Match{
		Board: board.NewBoard(matchId),
		Id:    matchId,
		P1:    &p1,
		P2:    &p2,
	}
	CurrentMatches.List[matchId] = &match
	match.Suggestion = &board.Position{X: board.SIZE / 2, Y: board.SIZE / 2}
	fmt.Println("New match started:", matchId)
	return &match
}

func (match *Match) AddMove(player *pl.Player, position *board.Position) error {
	err := match.CheckPlayersTurn(player)
	if err != nil {
		return err
	}
	err, toCapture := match.Board.PlaceStone(player, position, true)
	if err != nil {
		return err
	}
	move := &board.Move{player, position, toCapture}
	match.History = append(match.History, move)
	if match.Board.CheckWinningConditions(player, position) {
		match.Winner = player
	}
	match.Suggestion = heuristic.GetSuggestion(match.Board, move, GetOpponent(player))
	return nil
}

func (match *Match) CheckPlayersTurn(player *pl.Player) error {
	errMsg := ""
	if len(match.History) == 0 {
		if player.Id != 1 {
			errMsg = fmt.Sprintf("ERROR   (match %03d): It's not P%d's turn.\n", match.Id, player.Id)
		}
	} else if player.Id == match.History[len(match.History)-1].Player.Id {
		errMsg = fmt.Sprintf("ERROR   (match %03d): It's not P%d's turn.\n", match.Id, player.Id)
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func (match *Match) UnapplyLastMove() error {
	if len(match.History) == 0 {
		errMsg := fmt.Sprintf("ERROR   (match %03d): Match history is empty.\n", match.Id)
		return errors.New(errMsg)
	}
	// Save a reference to last move and remove it from history and board
	lastMove := match.History[len(match.History)-1]
	log.Printf("Unapplying move: %x\n", lastMove.Position)
	match.History = match.History[:len(match.History)-1]
	err := match.Board.RemoveStone(lastMove.Player, lastMove)
	if err != nil {
		return err
	}
	// Recalculate suggestion for player
	if len(match.History) == 0 {
		match.Suggestion = &board.Position{X: board.SIZE / 2, Y: board.SIZE / 2}
	} else {
		match.Suggestion = heuristic.GetSuggestion(match.Board, lastMove, lastMove.Player)
	}
	return nil
}

func PrintState(match *Match) {
	fmt.Printf("%29s\n", "-------------------")
	fmt.Printf(" %22s %d\n", "Match id:", match.Id)
	fmt.Printf("%29s\n", "-------------------")
	for _, line := range match.Board.Tab {
		fmt.Println(line)
	}
	fmt.Printf("%29s\n", "-------------------")
}
