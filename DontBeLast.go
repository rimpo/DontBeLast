package main

import (
	//	"container/list"
	"fmt"
)

func main() {
	/*var b Board
	b.Init()
	b.Print()
	b.PrintAllMoves()

	b.Move(4, 2, 1)

	b.Print()
	b.PrintAllMoves()

	b.Move(4, 3, 1)
	b.Print()
	b.PrintAllMoves()

	b.Move(0, 0, 1)
	b.Print()
	b.PrintAllMoves()
	*/

	var board Board
	var row, col, count int
	ai := AI{board: &board, myID: 0, opponentID: 1}
	board.Init()

	for {
		board.Print()
		fmt.Printf("Enter your moves(row, col, count):")
		fmt.Scanf("%d %d %d \n", &row, &col, &count)
		m := Move{Row: row, StartPos: col, Count: count}
		fmt.Printf("Your Move: %v\n", m)
		board.Move(m)
		m_ai := ai.EvaluateMove()
		board.Move(m_ai)
	}

}
