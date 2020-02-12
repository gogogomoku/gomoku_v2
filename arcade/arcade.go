package arcade

import (
	"fmt"

	"github.com/gogogomoku/gomoku_v2/board"
	pl "github.com/gogogomoku/gomoku_v2/player"
)

type Match struct {
	Board   *board.Board
	Id      int
	P1      *pl.Player
	P2      *pl.Player
	History []*Move
}

// A struct containing several simoultaneous matches
type Arcade struct {
	List    map[int]*Match
	Counter int
}

type Move struct {
	Player   *pl.Player
	Position *board.Position
}

// Global object containing a reference to simoultaneous matches
var CurrentMatches = Arcade{
	List:    make(map[int]*Match),
	Counter: 0,
}

// Creates a new match, stores it in Arcade map, returns it's address
func NewMatch() *Match {
	CurrentMatches.Counter++
	gameId := CurrentMatches.Counter
	p1 := pl.Player{Id: 1, Opponent: nil, Captured: 0}
	p2 := pl.Player{Id: 2, Opponent: &p1, Captured: 0}
	p1.Opponent = &p2
	match := Match{
		Board: board.NewBoard(gameId),
		Id:    gameId,
		P1:    &p1,
		P2:    &p2,
	}
	CurrentMatches.List[gameId] = &match
	fmt.Println("New game started:", gameId)
	return &match
}

func (match *Match) AddMove(player *pl.Player, position *board.Position) {
	match.Board.PlaceStone(player, position)
	match.History = append(match.History, &Move{player, position})
}

func PrintState(match *Match) {
	fmt.Printf("%29s\n", "-------------------")
	fmt.Printf(" %22s %d\n", "Game id:", match.Id)
	fmt.Printf("%29s\n", "-------------------")
	for _, line := range match.Board.Tab {
		fmt.Println(line)
	}
	fmt.Printf("%29s\n", "-------------------")
}
