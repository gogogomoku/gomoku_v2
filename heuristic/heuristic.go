package heuristic

import (
	"fmt"

	"github.com/gogogomoku/gomoku_v2/board"
	"github.com/gogogomoku/gomoku_v2/player"
)

func EvaluateBoard(b *board.Board, move *board.Position, player *player.Player) int {
	sequences := b.GetNSurroundingPositionsSequence(move, 6)
	maxSequence := 0
	for _, positionSequence := range *sequences {
		sequence := b.GetSequenceValues(&positionSequence)
		counter := 0
		for _, value := range *sequence {
			if value == player.Id {
				counter++
				if counter > maxSequence {
					maxSequence = counter
				}
			} else {
				counter = 0
			}
		}
	}
	return maxSequence * 100
}

func Suggest(b *board.Board, move *board.Position, player *player.Player) *board.Position {
	curY := move.Y - 5
	bestPosition := board.Position{X: -1, Y: -1}
	bestScore := -1
	for curY <= move.Y+5 {
		curX := move.X - 5
		for curX <= move.X+5 {
			position := board.Position{X: curX, Y: curY}
			if b.GetPositionValue(position) == 0 {
				tmpBoard := *b
				tmpBoard.PlaceStone(player, &position)
				fmt.Print(position)
				score := EvaluateBoard(&tmpBoard, &position, player)
				if score > bestScore {
					bestScore = score
					bestPosition = position
				}
			}
			curX++
		}
		curY++
	}
	fmt.Printf("bestPosition :: %#v\n", bestPosition)
	return &bestPosition
}
