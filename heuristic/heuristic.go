package heuristic

import (
	"github.com/gogogomoku/gomoku_v2/board"
)

func EvaluateBoard(b *board.Board, move *board.Position, playerId int8) int {
	sequences := b.GetNSurroundingPositionsSequence(move, 6)
	maxSequence := 0
	for _, positionSequence := range *sequences {
		sequence := b.GetSequenceValues(&positionSequence)
		counter := 0
		for _, value := range *sequence {
			if value == playerId {
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

func Suggest(b *board.Board, move *board.Position, playerId int8) {
	list := []board.Position{}
	startX := move.X - 2
	startY := move.Y - 2
	curX := startX
	curY := startY
	for curY <= startY+2 {
		curX = startX - 2
		for curX <= startX+2 {
			if b.Tab[curY][curX] == 0 {
				list = append(list, board.Position{X: curX, Y: curY})
			}
		}
		curY++
	}
}
