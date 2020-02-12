package board

import (
	"fmt"
	"github.com/gogogomoku/gomoku_v2/player"
	"reflect"
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

// Positions a stone in the board
func (b *Board) PlaceStone(player *player.Player, position *Position) {
	fmt.Printf("MOVE    (game %03d): player %d places at %x\n", b.GameId, player.Id, position)
	b.Tab[position.Y][position.X] = player.Id
	canCapture, toCapture := b.CheckCaptures(player, position)
	if canCapture {
		b.Capture(player, toCapture)
	}
}

// Check if by placing a stone, playerId can capture
// returns bool to true if can capturem and a slice of capturable positions
func (b *Board) CheckCaptures(player *player.Player, position *Position) (canCapture bool, toCapture *[]Position) {
	posToCapture := []Position{}
	pattern := [4]int8{player.Id, player.Opponent.Id, player.Opponent.Id, player.Id}
	for direction := int8(0); direction < 8; direction++ {

		// Build a sequence of positions to check capture
		tmpPosition := *position
		sequence := []Position{*position}
		for i := 0; i < 3; i++ {
			tmpPosition = b.GetNextPositionForDirection(tmpPosition, direction)
			if tmpPosition.X < 0 || tmpPosition.X > SIZE-1 {
				continue
			}
			if tmpPosition.Y < 0 || tmpPosition.Y > SIZE-1 {
				continue
			}
			sequence = append(sequence, tmpPosition)
		}
		if len(sequence) != 4 {
			continue
		}

		// Build a sequence of values from positions to check capture
		foundSequence := [4]int8{}
		for i := 0; i < 4; i++ {
			foundSequence[i] = b.GetPosition(sequence[i])
		}

		if reflect.DeepEqual(pattern, foundSequence) {
			canCapture = true
			posToCapture = append(posToCapture, sequence[1])
			posToCapture = append(posToCapture, sequence[2])
		}
	}
	return canCapture, &posToCapture
}

// Capture surrounding opponent stones for a playerId stone
func (b *Board) Capture(player *player.Player, toCapture *[]Position) {
	player.Captured += int8(len(*toCapture))
	for _, position := range *toCapture {
		b.Tab[position.Y][position.X] = 0
		fmt.Printf("CAPTURE (%03d): player %d captures %x\n", b.GameId, player.Id, position)
	}
}

func (b *Board) GetNextPositionForDirection(position Position, direction int8) Position {
	switch direction {
	case NW:
		return Position{position.X - 1, position.Y - 1}
	case N:
		return Position{position.X, position.Y - 1}
	case NE:
		return Position{position.X + 1, position.Y - 1}
	case W:
		return Position{position.X - 1, position.Y}
	case E:
		return Position{position.X + 1, position.Y}
	case SW:
		return Position{position.X - 1, position.Y + 1}
	case S:
		return Position{position.X, position.Y + 1}
	case SE:
		return Position{position.X + 1, position.Y + 1}
	}
	return Position{}
}

func (b *Board) GetPosition(position Position) int8 {
	return b.Tab[position.Y][position.X]
}
