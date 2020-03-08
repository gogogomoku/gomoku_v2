package player

type Player struct {
	Id         int8 `json:"id"`
	MatchId    int  `json:"matchId"`
	OpponentId int8 `json:"opponentId"`
	Captured   int8 `json:"captured"`
	IsAi       bool `json:"isAi"`
}
