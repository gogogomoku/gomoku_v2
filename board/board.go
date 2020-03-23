package board

import (
	"errors"
	"fmt"

	pl "github.com/gogogomoku/gomoku_v2/player"
)

// Needs to be bigger than 3 to allow captures
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
	Tab     [SIZE][SIZE]int8 `json:"tab"`
	MatchId int              `json:"matchId"`
}

type Position struct {
	X int8
	Y int8
}

func NewBoard(matchId int) *Board {
	b := Board{MatchId: matchId}
	return &b
}

// Places a stone in the board
func (b *Board) PlaceStone(player *pl.Player, position *Position, countCaptures bool) (err error, toCapture *[]Position) {
	if position.X < 0 || position.X >= SIZE || position.Y < 0 || position.Y >= SIZE {
		errMsg := fmt.Sprintf("ERROR   (match %03d): Position out of board. %x\n", b.MatchId, position)
		return errors.New(errMsg), nil
	}
	if b.Tab[position.Y][position.X] != 0 {
		errMsg := fmt.Sprintf("ERROR   (match %03d): Position is already occupied. %x\n", b.MatchId, position)
		return errors.New(errMsg), nil
	}
	b.Tab[position.Y][position.X] = player.Id
	canCapture, toCapture := b.CheckCaptures(player, position)
	if canCapture {
		b.Capture(player, toCapture, countCaptures)
	}
	return nil, toCapture
}

// Check if by placing a stone, playerId can capture
// returns bool to true if can capturem and a slice of capturable positions
func (b *Board) CheckCaptures(player *pl.Player, position *Position) (captures bool, list *[]Position) {
	// Store valid pattern for capture into an int32
	capturingPattern := int8ToInt32(
		[]int8{player.Id, player.OpponentId, player.OpponentId, player.Id},
	)
	posToCapture := []Position{}
	for direction := int8(0); direction < 8; direction++ {
		// Build a sequence of positions to check capture
		sequence := b.GetNPositionsSequence(position, direction, 4)
		if len(*sequence) != 4 {
			continue
		}
		sequenceValues := *b.GetSequenceValues(sequence)
		// First element changes to playerID to check if capture is possible even
		// In a board where the player hasn't placed stone yet
		sequenceValues[0] = player.Id
		// store found sequence into an int32 and compare
		foundPattern := int8ToInt32(sequenceValues)
		if foundPattern == capturingPattern {
			captures = true
			posToCapture = append(posToCapture, (*sequence)[1])
			posToCapture = append(posToCapture, (*sequence)[2])
		}
	}
	return captures, &posToCapture
}

// Capture surrounding opponent stones for a playerId stone
func (b *Board) Capture(player *pl.Player, toCapture *[]Position, countCaptures bool) {
	if countCaptures {
		player.Captured += int8(len(*toCapture))
	}
	for _, position := range *toCapture {
		b.Tab[position.Y][position.X] = 0
		fmt.Printf("CAPTURE (%03d): player %d captures %x\n", b.MatchId, player.Id, position)
	}
}

func (b *Board) CheckWinningConditions(player *pl.Player, position *Position) bool {
	return b.MoveCreatesFive(player, position)
}

// Check if placing a stone creates a winning sequence of 5 or more
func (b *Board) MoveCreatesFive(player *pl.Player, position *Position) bool {
	// Get bidirectional sequences to look for winning pattern
	sequences := b.GetNSurroundingPositionsSequence(position, 5)
	for _, seq := range *sequences {
		// For each sequence, count contiguous positions with player ID value
		counter := 0
		values := b.GetSequenceValues(&seq)
		for _, v := range *values {
			if v == player.Id {
				counter++
				if counter == 5 {
					return true
				}
			} else {
				counter = 0
			}
		}
	}
	return false
}
