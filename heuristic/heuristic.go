package heuristic

import (
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

func GetSuggestion(b *board.Board, lastMove *board.Position, player *player.Player) *board.Position {
	curY := lastMove.Y - 3
	bestPosition := board.Position{X: -1, Y: -1}
	bestScore := -1
	for curY <= lastMove.Y+3 {
		curX := lastMove.X - 3
		for curX <= lastMove.X+3 {
			position := board.Position{X: curX, Y: curY}
			if b.GetPositionValue(position) == 0 {
				tmpBoard := *b
				tmpBoard.PlaceStone(player, &position)
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
	// fmt.Printf("bestPosition :: %#v\n", bestPosition)
	return &bestPosition
}
