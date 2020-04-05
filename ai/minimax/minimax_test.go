package minimax

import (
	"math"
	"testing"

	"github.com/gogogomoku/gomoku_v2/ai/tree"
	"github.com/gogogomoku/gomoku_v2/arcade"
	"github.com/gogogomoku/gomoku_v2/arcade/match"
	"github.com/gogogomoku/gomoku_v2/board"
	pl "github.com/gogogomoku/gomoku_v2/player"
)

func evaluateBoardMock(b *board.Board, move *board.Position, player *pl.Player, fallback int) int {
	_ = player
	_ = move
	_ = b
	return fallback
}

func MakeNode(m *match.Match, playerId int8, score int) *tree.Node {
	player := m.P1
	if playerId == int8(2) {
		player = m.P2
	}
	children := []*tree.Node{}
	newNode := &tree.Node{
		CurrentBoard: m.Board,
		CurrentMove: &board.Move{
			Player:   player,
			Position: &board.Position{X: 0, Y: 0},
			Captures: &[]board.Position{},
		},
		CurrentPlayer: player,
		Children:      &children,
		Score:         score,
		BestChild:     &tree.Node{},
	}

	return newNode
}

func MakeTab(playerPositions *[]board.Position, opponentPositions *[]board.Position, playerId int8, matchId int) *[board.SIZE][board.SIZE]int8 {
	tab := board.NewBoard(matchId).Tab
	for _, pp := range *playerPositions {
		tab[pp.Y][pp.X] = playerId
	}
	for _, op := range *opponentPositions {
		tab[op.Y][op.X] = ^playerId & 0x3
	}
	return &tab
}

func makeSliceOfSlices(n uint, childScores ...[]int) [][]int {
	slices := make([][]int, n)
	for _, s := range childScores {
		slices = append(slices, s)
	}
	return slices
}

func makeTree(m *match.Match, node *tree.Node, nChildren int, depth int, scores *[]int) {
	if depth == 0 {
		node.Score = (*scores)[0]
		*scores = (*scores)[1:]
		return
	}
	for i := 0; i < nChildren; i++ {
		*node.Children = append(*node.Children, MakeNode(m, m.GetOpponent(node.CurrentPlayer).Id, 0))
		makeTree(m, (*node.Children)[i], nChildren, depth-1, scores)
	}
}

func TestMinimax_minimax(t *testing.T) {
	arcade.CurrentMatches = arcade.Arcade{
		List:    make(map[int]*match.Match),
		Counter: 0,
	}
	type fields struct {
		depth     int
		nChildren int
		scores    *[]int
		// nLeaves uint
		// leaves [][]int
		// parent *tree.Node
		// childScores [][][]int
	}
	tests := []struct {
		name      string
		wantScore int
		fields    fields
	}{
		{
			name:      "Depth 1, max score 1",
			wantScore: 1,
			fields: fields{
				depth:     1,
				nChildren: 1,
				scores:    &([]int{1}),
			},
		},
		{
			name:      "Depth 1, nChildren 4, max score 1",
			wantScore: 1,
			fields: fields{
				depth:     1,
				nChildren: 4,
				scores:    &([]int{1, -3, -3, -3}),
			},
		},
		{
			name:      "Depth 2, max score -3",
			wantScore: -3,
			fields: fields{
				depth:     2,
				nChildren: 2,
				scores:    &([]int{-4, -3, 1, -3}),
			},
		},
		{
			name:      "Depth 3, max score 0",
			wantScore: 0,
			fields: fields{
				depth:     3,
				nChildren: 2,
				scores:    &([]int{-4, -3, 1, -3, 0, 0, 0, 0}),
			},
		},
		{
			name:      "Depth 3, nChildren 3, want 8",
			wantScore: 8,
			fields: fields{
				depth:     3,
				nChildren: 3,
				scores: &([]int{
					-2, 5, 8, -4, 20, 15, 20, -5, 2,
					7, 2, 0, 50, 3, 15, 22, 11, 2,
					0, -1, -20, 12, 18, 3, 4, 2, 6,
				}),
			},
		},
		{
			name:      "Depth 5, nChildren 2, want 3",
			wantScore: 3,
			fields: fields{
				depth:     5,
				nChildren: 2,
				scores: &([]int{
					1, 2, 1, 2, 0, 0, -3, -1, 10, 9, 3, 5, 0, -10, -10, -11, -1, 1, 4, 3, 6, -6, 3, 2, 13, 10, 5, 7, -1, 2, 3, 16,
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newMatch := arcade.NewMatch(false, false)
			root := MakeNode(newMatch, 1, 0)
			makeTree(newMatch, root, tt.fields.nChildren, tt.fields.depth, tt.fields.scores)
			got := minimax(root, int8(tt.fields.depth), int(math.MinInt32), int(math.MaxInt32), true, evaluateBoardMock)
			if tt.wantScore != got {
				t.Errorf("got: %v, want: %v\n", got, tt.wantScore)
			}
			delete(arcade.CurrentMatches.List, arcade.CurrentMatches.Counter)
			arcade.CurrentMatches.Counter--
		})
	}
}
