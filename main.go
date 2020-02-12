package main

import (
	"fmt"
	"os"

	"github.com/gogogomoku/gomoku_v2/arcade"
	"github.com/gogogomoku/gomoku_v2/board"
	"github.com/gogogomoku/gomoku_v2/server"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser(
		"Gomoku",
		"Gomoku game server, for multiplayer and smart AI",
	)
	s := parser.Flag("s", "server",
		&argparse.Options{Help: "Run in server mode"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	if *s {
		server.StartServer()
	} else {
		StartLocalGomoku()
	}
}

func StartLocalGomoku() {
	fmt.Println("This multi-gomoku game thing is on!")
	game1 := arcade.NewMatch()
	game2 := arcade.NewMatch()
	game1.AddMove(game1.P1, &board.Position{X: 0, Y: 3})
	game1.AddMove(game1.P2, &board.Position{X: 0, Y: 2})
	game1.AddMove(game1.P1, &board.Position{X: 1, Y: 0})
	game1.AddMove(game1.P2, &board.Position{X: 0, Y: 1})
	game1.AddMove(game1.P1, &board.Position{X: 0, Y: 0})
	game2.AddMove(game1.P1, &board.Position{X: 0, Y: 3})
	game2.AddMove(game1.P2, &board.Position{X: 0, Y: 2})
	game2.AddMove(game1.P1, &board.Position{X: 0, Y: 0})
	game2.AddMove(game1.P2, &board.Position{X: 0, Y: 1})
	arcade.PrintState(game1)
	arcade.PrintState(game2)
}
