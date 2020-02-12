package board

import (
	"fmt"
	pl "github.com/gogogomoku/gomoku_v2/player"
)

const SIZE = 19

const (
	NW = iota
	N
	NE
	W
	E
	SW
	S
	SE
)

type Board struct {
	Tab    [SIZE][SIZE]int8
	GameId int
}

type Position struct {
	X int8
	Y int8
}

func NewBoard(gameId int) *Board {
	b := Board{GameId: gameId}
	return &b
}

// Places a stone in the board
func (b *Board) PlaceStone(player *pl.Player, position *Position) {
	fmt.Printf("MOVE    (game %03d): player %d places at %x\n", b.GameId, player.Id, position)
	b.Tab[position.Y][position.X] = player.Id
	canCapture, toCapture := b.CheckCaptures(player, position)
	if canCapture {
		b.Capture(player, toCapture)
	}
}

// Check if by placing a stone, playerId can capture
// returns bool to true if can capturem and a slice of capturable positions
func (b *Board) CheckCaptures(player *pl.Player, position *Position) (captures bool, list *[]Position) {
	// Store valid pattern for capture in 32 bits
	capturingPattern := int8IntoInt32(player.Id, player.Opponent.Id, player.Opponent.Id, player.Id)
	posToCapture := []Position{}
	for direction := int8(0); direction < 8; direction++ {
		// Build a sequence of positions to check capture
		sequence := b.GetNPositionsSequence(position, direction, 4)
		if len(sequence) != 4 {
			continue
		}
		// store found sequence in 32 bits and compare
		foundPattern := int8IntoInt32(
			b.GetPositionValue(sequence[0]),
			b.GetPositionValue(sequence[1]),
			b.GetPositionValue(sequence[2]),
			b.GetPositionValue(sequence[3]),
		)
		if foundPattern == capturingPattern {
			captures = true
			posToCapture = append(posToCapture, sequence[1])
			posToCapture = append(posToCapture, sequence[2])
		}
	}
	return captures, &posToCapture
}

// Capture surrounding opponent stones for a playerId stone
func (b *Board) Capture(player *pl.Player, toCapture *[]Position) {
	player.Captured += int8(len(*toCapture))
	for _, position := range *toCapture {
		b.Tab[position.Y][position.X] = 0
		fmt.Printf("CAPTURE (%03d): player %d captures %x\n", b.GameId, player.Id, position)
	}
}
