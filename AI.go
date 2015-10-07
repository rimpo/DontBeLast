package main

import (
	"fmt"
)

type AI struct {
	board      *Board
	myID       int
	opponentID int
}

func (ai *AI) Play(m Move, currentID int) int {
	ai.board.Move(m, currentID)
	//fmt.Printf("%d: play %v, (%d)\n", currentID, m, ai.board.CurrentCutCount)
	//ai.board.Print()
	moves := ai.board.GetAllMoves()
	defer ai.board.allocator.Release(moves)

	if ai.board.isLoser() {
		//fmt.Printf("%d: lost!!\n", currentID)
		ai.board.UndoMove(m)

		if currentID == ai.myID {
			return 0
		} else {
			return 100
		}
	}

	total := 0
	count := 0
	nextID := 0

	if currentID == ai.myID {
		nextID = ai.opponentID
	} else {
		nextID = ai.myID
	}

	for i := 0; i < moves.NoOfMoves; i++ {
		if ai.board.isValidMove(moves.AllMoves[i]) {

			total += ai.Play(moves.AllMoves[i], nextID)
			count++
		}
	}
	ai.board.UndoMove(m)
	if total > 0 {
		return total / count
	} else {
		return 0
	}
}

func (ai *AI) EvaluateMove() Move {

	moves := ai.board.GetAllMoves()
	defer ai.board.allocator.Release(moves)
	//fmt.Printf("total moves: %d\n", moves.NoOfMoves)

	scores := make([]int, moves.NoOfMoves)
	var maxScore, bestMove int

	for i := 0; i < moves.NoOfMoves; i++ {
		if ai.board.isValidMove(moves.AllMoves[i]) {
			scores[i] = ai.Play(moves.AllMoves[i], ai.myID)
			if maxScore < scores[i] {
				maxScore = scores[i]
				bestMove = i
			}
			break
		}
	}
	fmt.Printf("scores: %v\n", scores)
	return moves.AllMoves[bestMove]
}
