package arcade

import (
	"fmt"

	"github.com/gogogomoku/gomoku_v2/arcade/match"
	"github.com/gogogomoku/gomoku_v2/board"
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
	newMatch := match.CreateMatch(aiP1, aiP2, matchId)
	CurrentMatches.List[matchId] = newMatch
	newMatch.Suggestion = &board.Position{X: board.SIZE / 2, Y: board.SIZE / 2}
	fmt.Println("New match started:", matchId)
	return newMatch
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
