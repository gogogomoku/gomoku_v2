package ai

import (
	"log"

	"github.com/gogogomoku/gomoku_v2/board"
	"github.com/gogogomoku/gomoku_v2/player"
)

type Node struct {
	CurrentBoard  *board.Board
	CurrentMove   *board.Move
	CurrentPlayer *player.Player
	Children      *[]*Node
	Score         int   // for evalution
	BestChild     *Node // for evaluation
}

func GetPossibleMoves(b *board.Board, lastMove *board.Position, player *player.Player) *[]*board.Move {
	curY := lastMove.Y - 1
	possibleMoves := []*board.Move{}
	for curY <= lastMove.Y+1 {
		curX := lastMove.X - 1
		for curX <= lastMove.X+1 {
			position := board.Position{X: curX, Y: curY}
			if b.GetPositionValue(position) == 0 {
				_, captures := b.CheckCaptures(player, &position)
				possibleMoves = append(possibleMoves, &board.Move{Player: player, Position: &position, Captures: captures})
			}
			curX++
		}
		curY++
	}
	return &possibleMoves
}

func projectPossibleMoves(parent *Node, maxDepth int8) {
	if maxDepth == 0 {
		return
	}
	possiblePositions := GetPossibleMoves(parent.CurrentBoard, parent.CurrentMove.Position, parent.CurrentPlayer)
	for _, move := range *possiblePositions {
		*parent.Children = append(*parent.Children, &Node{CurrentBoard: parent.CurrentBoard, CurrentMove: move, CurrentPlayer: parent.CurrentMove.Player, Children: &[]*Node{}, Score: 0, BestChild: nil})
	}
	for _, child := range *parent.Children {
		err, _ := child.CurrentBoard.PlaceStone(child.CurrentMove.Player, child.CurrentMove.Position, true)
		if err != nil {
			log.Printf("Error in projectPossibleMoves (PlaceStone): %v", err)
			return
		}
		projectPossibleMoves(child, maxDepth-1)
		err = child.CurrentBoard.RemoveStone(child.CurrentMove.Player, child.CurrentMove)
		if err != nil {
			log.Printf("Error in projectPossibleMoves (RemoveStone): %v", err)
			return
		}
	}
}

func EvaluateBoard(b *board.Board, move *board.Position, player *player.Player) int {
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
	return maxSequence * 100
}

func getDumbBestPosition(b *board.Board, lastMove *board.Position, player *player.Player) *board.Position {
	curY := lastMove.Y - 3
	bestPosition := board.Position{X: -1, Y: -1}
	bestScore := -1
	for curY <= lastMove.Y+3 {
		curX := lastMove.X - 3
		for curX <= lastMove.X+3 {
			position := board.Position{X: curX, Y: curY}
			if b.GetPositionValue(position) == 0 {
				tmpBoard := *b
				tmpBoard.PlaceStone(player, &position, false)
				score := EvaluateBoard(&tmpBoard, &position, player)
				if score > bestScore {
					bestScore = score
					bestPosition = position
				}
			}
			curX++
		}
		curY++
	}
	// fmt.Printf("bestDumbPosition :: %#v\n", bestDumbPosition)
	return &bestPosition

}

func GetSuggestion(b *board.Board, lastMove *board.Move, nextPlayer *player.Player) *board.Position {
	boardCopy := &board.Board{Tab: b.Tab, MatchId: b.MatchId}
	nextPlayerCopy := *nextPlayer
	lastPlayerCopy := *lastMove.Player
	moveCopy := board.Move{Player: &lastPlayerCopy, Position: lastMove.Position, Captures: lastMove.Captures}
	state := &Node{CurrentBoard: boardCopy, CurrentMove: &moveCopy, CurrentPlayer: &nextPlayerCopy, Children: &[]*Node{}, Score: 0, BestChild: nil}

	projectPossibleMoves(state, 1)
	bestPosition := getDumbBestPosition(b, lastMove.Position, &nextPlayerCopy)
	return bestPosition
}
