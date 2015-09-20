package main

import (
//	"container/list"
//	"fmt"
)

func main() {
	var b Board
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

}
