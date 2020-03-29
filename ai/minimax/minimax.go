package minimax

import (
	"fmt"
	"math"

	"github.com/gogogomoku/gomoku_v2/ai/tree"
	"github.com/gogogomoku/gomoku_v2/board"
	pl "github.com/gogogomoku/gomoku_v2/player"
)

func LaunchMinimax(root *tree.Node, depth int8) {
	fmt.Printf("Minimax with Depth %d\n", depth)
	value := minimax(root, depth, int(math.MinInt32), int(math.MaxInt32), true)
	fmt.Printf("Final value %d\n", value)
}

func minimax(node *tree.Node, depth int8, alpha int, beta int, maximize bool) int {
	// // TODO: Add if state is a win to this condition
	if depth == 0 {
		return EvaluateBoard(node.CurrentBoard, node.CurrentMove.Position, node.CurrentPlayer)
	}

	if maximize {
		maxEval := int(math.MinInt32)
		for _, child := range *node.Children {
			eval := minimax(child, depth-1, alpha, beta, false)
			maxEval = maximum(maxEval, eval)
			alpha = maximum(alpha, eval)
			if node.BestChild == nil || child.Score > node.BestChild.Score {
				node.BestChild = child
			}
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		minEval := int(math.MaxInt32)
		for _, child := range *node.Children {
			eval := minimax(child, depth-1, alpha, beta, true)
			minEval = minimum(minEval, eval)
			beta = minimum(beta, eval)
			if node.BestChild == nil || child.Score < node.BestChild.Score {
				node.BestChild = child
			}
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
	return 0
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func EvaluateBoard(b *board.Board, move *board.Position, player *pl.Player) int {
	sequences := b.GetNSurroundingPositionsSequence(move, 6)
	maxSequence := 0
	for _, positionSequence := range *sequences {
		sequence := b.GetSequenceValues(&positionSequence)
		counter := 0
		for _, value := range *sequence {
			if value == player.Id {
				counter++
				if counter > maxSequence {
					maxSequence = counter
				}
			} else {
				counter = 0
			}
		}
	}
	return int(math.Pow(10, float64(maxSequence)))
}
