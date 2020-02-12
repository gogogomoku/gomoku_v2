package board

import (
	"github.com/gogogomoku/gomoku_v2/player"
)

const SIZE = 19

type Board [19][19]int8

type GomokuBoard interface {
	PlaceStone(playerId int8, posX int8, posY int8)
	CheckCaptures(posX int8, posY int8)
	Capture(posX int8, posY int8)
}

func NewBoard() *Board {
	b := Board{}
	return &b
}

// Positions a stone in the board
func (b *Board) PlaceStone(player *player.Player, posX int8, posY int8) {
	b[posY][posX] = player.Id
}

// Check if by placing a stone, playerId can capture
func (b *Board) CheckCaptures(player *player.Player, posX int8, posY int8) {

}

// Capture surrounding opponent stones for a playerId stone
func (b *Board) Capture(player *player.Player, posX int8, posY int8) {

}
