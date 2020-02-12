package player

type Player struct {
	Id       int8
	Opponent *Player
	Captured int8
}
