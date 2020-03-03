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
	match1 := arcade.NewMatch(true, true)
	match2 := arcade.NewMatch(true, true)
	match1.AddMove(match1.P1, &board.Position{X: 0, Y: 3})
	match1.AddMove(match1.P2, &board.Position{X: 0, Y: 2})
	match1.AddMove(match1.P1, &board.Position{X: 1, Y: 0})
	match1.AddMove(match1.P2, &board.Position{X: 0, Y: 1})
	match1.AddMove(match1.P1, &board.Position{X: 0, Y: 0})
	match2.AddMove(match1.P1, &board.Position{X: 0, Y: 3})
	match2.AddMove(match1.P2, &board.Position{X: 0, Y: 2})
	match2.AddMove(match1.P1, &board.Position{X: 0, Y: 0})
	match2.AddMove(match1.P2, &board.Position{X: 0, Y: 1})
	arcade.PrintState(match1)
	arcade.PrintState(match2)
}
