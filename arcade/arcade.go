package arcade

import (
	"fmt"
	"github.com/gogogomoku/gomoku_v2/board"
	"github.com/gogogomoku/gomoku_v2/player"
	"time"
)

type Match struct {
	Board *board.Board
	Id    string
	P1    *player.Player
	P2    *player.Player
}

// A struct containing several simoultaneous matches
type Arcade struct {
	List map[string]*Match
}

// Global object containing a reference to simoultaneous matches
var CurrentMatches = Arcade{
	List: make(map[string]*Match),
}

// Creates a new match, stores it in Arcade map, returns it's address
func NewMatch() *Match {
	gameId := fmt.Sprint(time.Now().UnixNano())
	match := Match{
		Board: board.NewBoard(),
		Id:    gameId,
		P1:    &player.Player{Id: 1},
		P2:    &player.Player{Id: 2},
	}
	CurrentMatches.List[gameId] = &match
	fmt.Println("New game started:", gameId)
	return &match
}

func PrintState(match *Match) {
	fmt.Printf("%29s\n", "-------------------")
	fmt.Printf("%29s   \n", match.Id)
	fmt.Printf("%29s\n", "-------------------")
	for _, line := range match.Board {
		fmt.Println(line)
	}
	fmt.Printf("%29s\n", "-------------------")
}
