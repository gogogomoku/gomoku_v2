package minimax

import (
	"testing"

	"github.com/gogogomoku/gomoku_v2/ai/tree"
)

func TestMinimax_LaunchMinimax(t *testing.T) {
	type args struct {
		root  *tree.Node
		depth int8
	}
	tests := []struct {
		name string
		args args
		// wantCaptures bool
	}{
		{
			name: "Basic minimax test",
			args: args{
				root:  &tree.Node{},
				depth: 3,
			},
			// wantCaptures: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if gotCaptures != tt.wantCaptures {
			// t.Errorf("Board.CheckCaptures() gotCaptures = %v, want %v", gotCaptures, tt.wantCaptures)
			// }
		})
	}
}
