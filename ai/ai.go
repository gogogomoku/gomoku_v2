package ai

import (
	"github.com/gogogomoku/gomoku_v2/ai/minimax"
	"github.com/gogogomoku/gomoku_v2/ai/tree"
	"github.com/gogogomoku/gomoku_v2/board"
	"github.com/gogogomoku/gomoku_v2/player"
)

func GetSuggestion(b *board.Board, lastMove *board.Move, nextPlayer *player.Player) *board.Position {
	boardCopy := &board.Board{Tab: b.Tab, MatchId: b.MatchId}
	nextPlayerCopy := *nextPlayer
	lastPlayerCopy := *lastMove.Player
	moveCopy := board.Move{Player: &lastPlayerCopy, Position: lastMove.Position, Captures: lastMove.Captures}
	state := &tree.Node{CurrentBoard: boardCopy, CurrentMove: &moveCopy, CurrentPlayer: &nextPlayerCopy, Children: &[]*tree.Node{}, Score: 0, BestChild: nil}

	depth := int8(4)
	tree.ProjectPossibleMoves(state, depth)
	minimax.LaunchMinimax(state, depth)
	bestPosition := state.BestChild.CurrentMove.Position
	return bestPosition
}
