package arcade

import (
	"fmt"

	"github.com/gogogomoku/gomoku_v2/arcade/match"
	"github.com/gogogomoku/gomoku_v2/board"
	pl "github.com/gogogomoku/gomoku_v2/player"
)

// A struct containing several simoultaneous matches
type Arcade struct {
	List    map[int]*match.Match
	Counter int
}

// Global object containing a reference to simoultaneous matches
var CurrentMatches = Arcade{
	List:    make(map[int]*match.Match),
	Counter: 0,
}

// Creates a new match, stores it in Arcade map, returns it's address
func NewMatch(aiP1 bool, aiP2 bool) *match.Match {
	CurrentMatches.Counter++
	matchId := CurrentMatches.Counter
	p1 := pl.Player{Id: 1, OpponentId: 2, Captured: 0, IsAi: aiP1, MatchId: matchId}
	p2 := pl.Player{Id: 2, OpponentId: 1, Captured: 0, IsAi: aiP2, MatchId: matchId}
	match := match.Match{
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

func PrintState(match *match.Match) {
	fmt.Printf("%29s\n", "-------------------")
	fmt.Printf(" %22s %d\n", "Match id:", match.Id)
	fmt.Printf("%29s\n", "-------------------")
	for _, line := range match.Board.Tab {
		fmt.Println(line)
	}
	fmt.Printf("%29s\n", "-------------------")
}
