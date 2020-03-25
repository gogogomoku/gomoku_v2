package arcade

import (
	"reflect"
	"testing"

	"github.com/gogogomoku/gomoku_v2/board"
	pl "github.com/gogogomoku/gomoku_v2/player"
)

func MakeBoard(playerPositions *[]board.Position, opponentPositions *[]board.Position, playerId int8, matchId int) *[board.SIZE][board.SIZE]int8 {
	tab := board.NewBoard(matchId).Tab
	for _, pp := range *playerPositions {
		tab[pp.Y][pp.X] = playerId
	}
	for _, op := range *opponentPositions {
		tab[op.Y][op.X] = ^playerId & 0x3
	}
	return &tab
}

func TestArcade_GetOpponent(t *testing.T) {
	type fields struct {
		Match *Match
	}
	type args struct {
		player *pl.Player
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantPlayer *pl.Player
	}{
		{
			name: "Get opponent of player 1",
			fields: fields{
				Match: NewMatch(false, false),
			},
			args: args{
				player: CurrentMatches.List[1].P1,
			},
			wantPlayer: CurrentMatches.List[1].P2,
		},
		{
			name: "Get opponent of player 2",
			fields: fields{
				Match: NewMatch(false, false),
			},
			args: args{
				player: CurrentMatches.List[2].P2,
			},
			wantPlayer: CurrentMatches.List[2].P1,
		},
		{
			name: "Get opponent of nil player pointer",
			fields: fields{
				Match: NewMatch(false, false),
			},
			args: args{
				player: nil,
			},
			wantPlayer: nil,
		},
	}

	for _, tt := range tests {
		gotPlayer := GetOpponent(tt.args.player)
		t.Run(tt.name, func(t *testing.T) {
			if GetOpponent(tt.args.player) != tt.wantPlayer {
				t.Errorf("Arcade.GetOpponent() got Player = %v, want %v", gotPlayer, tt.wantPlayer)
			}
		})
	}
}

func TestArcade_NewMatch(t *testing.T) {
	type args struct {
		aiP1 bool
		aiP2 bool
	}
	tests := []struct {
		name      string
		args      args
		wantMatch *Match
	}{
		{
			name: "New match no AI",
			args: args{
				aiP1: false,
				aiP2: false,
			},
			wantMatch: &Match{
				Board: board.NewBoard(CurrentMatches.Counter + 1),
				Id:    CurrentMatches.Counter + 1,
				P1:    &pl.Player{Id: 1, OpponentId: 2, Captured: 0, IsAi: false, MatchId: CurrentMatches.Counter + 1},
				P2:    &pl.Player{Id: 2, OpponentId: 1, Captured: 0, IsAi: false, MatchId: CurrentMatches.Counter + 1},
			},
		},
		{
			name: "New match both AI",
			args: args{
				aiP1: true,
				aiP2: true,
			},
			wantMatch: &Match{
				Board: board.NewBoard(CurrentMatches.Counter + 2),
				Id:    CurrentMatches.Counter + 2,
				P1:    &pl.Player{Id: 1, OpponentId: 2, Captured: 0, IsAi: true, MatchId: CurrentMatches.Counter + 2},
				P2:    &pl.Player{Id: 2, OpponentId: 1, Captured: 0, IsAi: true, MatchId: CurrentMatches.Counter + 2},
			},
		},
	}

	for _, tt := range tests {
		gotMatch := NewMatch(tt.args.aiP1, tt.args.aiP2)
		t.Run(tt.name, func(t *testing.T) {
			if gotMatch.Id != tt.wantMatch.Id {
				t.Errorf("Arcade.NewMatch() got match.Id = %v, want %v", gotMatch.Id, tt.wantMatch.Id)
			}
			if gotMatch.Board.MatchId != tt.wantMatch.Board.MatchId {
				t.Errorf("Arcade.NewMatch() got match.Board.Id = %v, want %v", gotMatch.Board.MatchId, tt.wantMatch.Board.MatchId)
			}
			if gotMatch.P1.IsAi != tt.wantMatch.P1.IsAi {
				t.Errorf("Arcade.NewMatch() got match.P1.IsAi = %v, want %v", gotMatch.P1.IsAi, tt.wantMatch.P1.IsAi)
			}
			if gotMatch.P2.IsAi != tt.wantMatch.P2.IsAi {
				t.Errorf("Arcade.NewMatch() got match.P2.IsAi = %v, want %v", gotMatch.P2.IsAi, tt.wantMatch.P2.IsAi)
			}
		})
	}
}

func TestArcade_AddMove(t *testing.T) {
	type testMove struct {
		playerId int8
		position *board.Position
	}
	type args struct {
		moves []*testMove
	}
	tests := []struct {
		name    string
		args    args
		wantTab *[board.SIZE][board.SIZE]int8
	}{
		{
			name: "Place simple stone P1 empty board",
			args: args{
				moves: []*testMove{
					&testMove{
						playerId: 1,
						position: &board.Position{X: 0, Y: 0},
					},
				},
			},
			wantTab: MakeBoard(
				&[]board.Position{
					board.Position{0, 0},
				},
				&[]board.Position{},
				int8(1),
				CurrentMatches.Counter+1,
			),
		},
		{
			name: "Place P1 {0,0} P2 {1,0}",
			args: args{
				moves: []*testMove{
					&testMove{
						playerId: 1,
						position: &board.Position{X: 0, Y: 0},
					},
					&testMove{
						playerId: 2,
						position: &board.Position{X: 1, Y: 0},
					},
				},
			},
			wantTab: MakeBoard(
				&[]board.Position{
					board.Position{0, 0},
				},
				&[]board.Position{
					board.Position{1, 0},
				},
				int8(1),
				CurrentMatches.Counter+1,
			),
		},
		{
			name: "Capture: Place P1 {0,0} P2 {1,0} P1 {3,3} P2 {2,0} P1 {3,0}",
			args: args{
				moves: []*testMove{
					&testMove{
						playerId: 1,
						position: &board.Position{X: 0, Y: 0},
					},
					&testMove{
						playerId: 2,
						position: &board.Position{X: 1, Y: 0},
					},
					&testMove{
						playerId: 1,
						position: &board.Position{X: 3, Y: 3},
					},
					&testMove{
						playerId: 2,
						position: &board.Position{X: 2, Y: 0},
					},
					&testMove{
						playerId: 1,
						position: &board.Position{X: 3, Y: 0},
					},
				},
			},
			wantTab: MakeBoard(
				&[]board.Position{
					board.Position{0, 0},
					board.Position{3, 3},
					board.Position{3, 0},
				},
				&[]board.Position{},
				int8(1),
				CurrentMatches.Counter+1,
			),
		},
		{
			name: "Error case: Place simple stone P2 empty board",
			args: args{
				moves: []*testMove{
					&testMove{
						playerId: 2,
						position: &board.Position{X: 0, Y: 0},
					},
				},
			},
			wantTab: MakeBoard(
				&[]board.Position{},
				&[]board.Position{},
				int8(1),
				CurrentMatches.Counter+1,
			),
		},
		{
			name: "Error case: Place p1 outside board",
			args: args{
				moves: []*testMove{
					&testMove{
						playerId: 1,
						position: &board.Position{X: -1, Y: -1},
					},
				},
			},
			wantTab: MakeBoard(
				&[]board.Position{},
				&[]board.Position{},
				int8(1),
				CurrentMatches.Counter+1,
			),
		},
	}

	for _, tt := range tests {
		gotMatch := NewMatch(false, false)
		t.Run(tt.name, func(t *testing.T) {
			for _, m := range tt.args.moves {
				if m.playerId == 1 {
					gotMatch.AddMove(gotMatch.P1, m.position)
				}
				if m.playerId == 2 {
					gotMatch.AddMove(gotMatch.P2, m.position)
				}
			}
			if !reflect.DeepEqual(gotMatch.Board.Tab, *tt.wantTab) {
				t.Errorf("Arcade.AddMove() got match.Board.Tab = %v, want %v", gotMatch.Board.Tab, tt.wantTab)
			}
		})
	}
}
