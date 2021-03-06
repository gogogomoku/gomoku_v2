package board

import (
	"reflect"
	"testing"

	pl "github.com/gogogomoku/gomoku_v2/player"
)

func MakeBoard(playerPositions *[]Position, opponentPositions *[]Position, playerId int8) *[SIZE][SIZE]int8 {
	tab := NewBoard(0).Tab
	for _, pp := range *playerPositions {
		tab[pp.Y][pp.X] = playerId
	}
	for _, op := range *opponentPositions {
		tab[op.Y][op.X] = ^playerId & 0x3
	}
	return &tab
}

func TestBoard_CheckCaptures(t *testing.T) {
	type fields struct {
		Tab [SIZE][SIZE]int8
	}
	type args struct {
		player   *pl.Player
		position *Position
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCaptures bool
		wantList     *[]Position
	}{
		{
			name: "Nothing to capture (no stones)",
			fields: fields{
				Tab: NewBoard(0).Tab,
			},
			args: args{
				player: &pl.Player{
					Id:         1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: 0, Y: 0},
			},
			wantCaptures: false,
			wantList:     &[]Position{},
		},
		{
			name: "Still nothing to capture (current player stones)",
			fields: fields{
				Tab: *MakeBoard(
					&[]Position{
						Position{0, 0},
						Position{1, 0},
						Position{2, 0},
					},
					&[]Position{},
					int8(1)),
			},
			args: args{
				player: &pl.Player{Id: 1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: 3, Y: 0},
			},
			wantCaptures: false,
			wantList:     &[]Position{},
		},
		{
			name: "Capture W->E",
			fields: fields{
				Tab: *MakeBoard(
					&[]Position{
						Position{0, 0},
					},
					&[]Position{
						Position{1, 0},
						Position{2, 0},
					},
					int8(1)),
			},
			args: args{
				player: &pl.Player{
					Id:         1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: 3, Y: 0},
			},
			wantCaptures: true,
			wantList:     &[]Position{Position{2, 0}, Position{1, 0}},
		},
		{
			name: "Capture NW->SE",
			fields: fields{
				Tab: *MakeBoard(
					&[]Position{
						Position{0, 0},
					},
					&[]Position{
						Position{1, 1},
						Position{2, 2},
					},
					int8(1)),
			},
			args: args{
				player: &pl.Player{
					Id:         1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: 3, Y: 3},
			},
			wantCaptures: true,
			wantList:     &[]Position{Position{2, 2}, Position{1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Tab: tt.fields.Tab,
			}
			gotCaptures, gotList := b.CheckCaptures(tt.args.player, tt.args.position)
			if gotCaptures != tt.wantCaptures {
				t.Errorf("Board.CheckCaptures() gotCaptures = %v, want %v", gotCaptures, tt.wantCaptures)
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("Board.CheckCaptures() gotList = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}

func TestBoard_PlaceStone(t *testing.T) {
	type fields struct {
		Tab     [SIZE][SIZE]int8
		MatchId int
	}
	type args struct {
		player   *pl.Player
		position *Position
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue int8
	}{
		{
			name: "Place at (0, 0)",
			fields: fields{
				Tab:     NewBoard(0).Tab,
				MatchId: 1,
			},
			args: args{
				player: &pl.Player{
					Id:         1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: 0, Y: 0},
			},
			wantValue: 1,
		},
		{
			name: "Place at (SIZE-1, SIZE-1)",
			fields: fields{
				Tab:     NewBoard(0).Tab,
				MatchId: 1,
			},
			args: args{
				player: &pl.Player{Id: 1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: SIZE - 1, Y: SIZE - 1},
			},
			wantValue: 1,
		},
		{
			name: "Place at (-1, 0)",
			fields: fields{
				Tab:     NewBoard(0).Tab,
				MatchId: 1,
			},
			args: args{
				player: &pl.Player{
					Id:         1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: -1, Y: 0},
			},
			wantValue: -1,
		},
		{
			name: "Place at (0, -1)",
			fields: fields{
				Tab:     NewBoard(0).Tab,
				MatchId: 1,
			},
			args: args{
				player: &pl.Player{
					Id:         1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: 0, Y: -1},
			},
			wantValue: -1,
		},
		{
			name: "Place at (SIZE, 0)",
			fields: fields{
				Tab:     NewBoard(0).Tab,
				MatchId: 1,
			},
			args: args{
				player: &pl.Player{
					Id:         1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: SIZE, Y: 0},
			},
			wantValue: -1,
		},
		{
			name: "Place at (0, SIZE)",
			fields: fields{
				Tab:     NewBoard(0).Tab,
				MatchId: 1,
			},
			args: args{
				player: &pl.Player{
					Id:         1,
					OpponentId: 2,
					Captured:   0,
					IsAi:       false,
				},
				position: &Position{X: 0, Y: SIZE},
			},
			wantValue: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Tab:     tt.fields.Tab,
				MatchId: tt.fields.MatchId,
			}
			b.PlaceStone(tt.args.player, tt.args.position, true)
			gotValue := b.GetPositionValue(*tt.args.position)
			if gotValue != tt.wantValue {
				t.Errorf("Board.CheckCaptures() gotCaptures = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
