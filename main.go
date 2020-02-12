package main

import (
	"fmt"
	"github.com/gogogomoku/gomoku_v2/arcade"
)

func main() {
	fmt.Println("This multi-gomoku game thing is on!")
	g1 := arcade.NewMatch()
	g2 := arcade.NewMatch()
	g1.Board.PlaceStone(g1.P1, 1, 1)
	g1.Board.PlaceStone(g1.P2, 1, 2)
	g2.Board.PlaceStone(g1.P2, 1, 1)
	g2.Board.PlaceStone(g1.P1, 1, 2)
	arcade.PrintState(g1)
	arcade.PrintState(g2)
}
