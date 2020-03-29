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
	for _, ch1 := range *root.Children {
		fmt.Println("****------****")
		fmt.Println()
		fmt.Println(*ch1.CurrentMove.Position, ch1.Score)
		fmt.Println()
		// for _, ch2 := range *ch1.Children {
		// 	fmt.Println(*ch2.CurrentMove.Position, ch2.Score)
		// }
	}
}

func minimax(node *tree.Node, depth int8, alpha int, beta int, maximize bool) int {
	// // TODO: Add if state is a win to this condition
	// PrintState(node.CurrentBoard)
	if depth == 0 || len(*node.Children) == 0 {
		node.Score = EvaluateBoard(node.CurrentBoard, node.CurrentMove.Position, node.CurrentMove.Player)
		return node.Score
	}

	if maximize {
		maxEval := int(math.MinInt32)
		for _, child := range *node.Children {
			child.CurrentBoard.PlaceStone(child.CurrentMove.Player, child.CurrentMove.Position, true)
			eval := minimax(child, depth-1, alpha, beta, false)
			child.CurrentBoard.RemoveStone(child.CurrentMove.Player, child.CurrentMove)
			maxEval = maximum(maxEval, eval)
			alpha = maximum(alpha, eval)
			if node.BestChild == nil || child.Score > node.BestChild.Score {
				node.BestChild = child
				node.Score = node.BestChild.Score
			}
			if beta <= alpha {
				// fmt.Println("PRUNING")
				break
			}
		}
		return maxEval
	} else {
		minEval := int(math.MaxInt32)
		for _, child := range *node.Children {
			child.CurrentBoard.PlaceStone(child.CurrentMove.Player, child.CurrentMove.Position, true)
			eval := minimax(child, depth-1, alpha, beta, true)
			child.CurrentBoard.RemoveStone(child.CurrentMove.Player, child.CurrentMove)
			minEval = minimum(minEval, eval)
			beta = minimum(beta, eval)
			if node.BestChild == nil || child.Score < node.BestChild.Score {
				node.BestChild = child
				node.Score = node.BestChild.Score
			}
			if beta <= alpha {
				// fmt.Println("PRUNING")
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
	sequences := [][board.SIZE]int8{}
	// Add horizontal sequences
	for _, line := range b.Tab {
		sequences = append(sequences, line)
	}
	// Add vertical sequences
	for i := 0; i < board.SIZE; i++ {
		verticalSequence := [board.SIZE]int8{}
		for line := 0; line < board.SIZE; line++ {
			verticalSequence[line] = b.Tab[line][i]
		}
		sequences = append(sequences, verticalSequence)
	}
	// fmt.Println(sequences)
	maxSequence := 0
	playerValue := 0
	for _, sequence := range sequences {
		counter := 0
		for _, value := range sequence {
			if value == player.Id {
				counter++
				if counter > maxSequence {
					maxSequence = counter
				}
			} else {
				counter = 0
			}
		}
		playerValue += int(math.Pow(10, float64(maxSequence)))
	}
	maxSequenceOpponent := 0
	opponentValue := 0
	for _, sequence := range sequences {
		counter := 0
		for _, value := range sequence {
			if value == player.OpponentId {
				counter++
				if counter > maxSequenceOpponent {
					maxSequenceOpponent = counter
				}
			} else {
				counter = 0
			}
		}
		opponentValue += int(math.Pow(10, float64(maxSequenceOpponent)))
	}

	finalValue := playerValue - int(1.2*float64(opponentValue))
	// fmt.Println("playerValue: ", playerValue)
	// fmt.Println("opponentValue: ", opponentValue)
	// fmt.Println("finalValue: ", finalValue)
	return finalValue
}

func PrintState(board *board.Board) {
	fmt.Printf("%29s\n", "-------------------")
	fmt.Printf(" %22s %d\n", "Match id:", board.MatchId)
	fmt.Printf("%29s\n", "-------------------")
	for _, line := range board.Tab {
		fmt.Println(line)
	}
	fmt.Printf("%29s\n", "-------------------")
}
