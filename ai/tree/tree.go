package tree

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
	curY := lastMove.Y - 8
	possibleMoves := []*board.Move{}
	for curY <= lastMove.Y+8 {
		curX := lastMove.X - 8
		for curX <= lastMove.X+8 {
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

func ProjectPossibleMoves(parent *Node, maxDepth int8) {
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
		ProjectPossibleMoves(child, maxDepth-1)
		err = child.CurrentBoard.RemoveStone(child.CurrentMove.Player, child.CurrentMove)
		if err != nil {
			log.Printf("Error in projectPossibleMoves (RemoveStone): %v", err)
			return
		}
	}
}
