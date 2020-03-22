package arcade

import (
	"fmt"

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
	History    []*Move         `json:"history"`
}

// A struct containing several simoultaneous matches
type Arcade struct {
	List    map[int]*Match
	Counter int
}

type Move struct {
	Player   *pl.Player        `json:"player"`
	Position *board.Position   `json:"position"`
	Captures *[]board.Position `json:"captures"`
}

// Global object containing a reference to simoultaneous matches
var CurrentMatches = Arcade{
	List:    make(map[int]*Match),
	Counter: 0,
}

func GetOpponent(player *player.Player) *player.Player {
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
	err, toCapture := match.Board.PlaceStone(player, position, true)
	if err != nil {
		return err
	}
	match.History = append(match.History, &Move{player, position, toCapture})
	if match.Board.CheckWinningConditions(player, position) {
		match.Winner = player
	}
	match.Suggestion = heuristic.GetSuggestion(match.Board, position, GetOpponent(player))
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
