package player

type Player struct {
	Id         int8 `json:"id"`
	OpponentId int8 `json:"OpponentId"`
	Captured   int8 `json:"captured"`
}
