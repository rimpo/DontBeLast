package main

import (
	//	"container/list"
	"flag"
	//"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
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
	//var row, col, count int
	ai := AI{board: &board, myID: 1, opponentID: 2}
	board.Init()

	for {
		//board.Print()
		//fmt.Printf("Enter your moves(row, col, count):")
		//fmt.Scanf("%d %d %d \n", &row, &col, &count)
		//m := Move{Row: row, StartPos: col, Count: count}
		m := Move{Row: 4, StartPos: 2, Count: 1}
		//fmt.Printf("Your Move: %v\n", m)
		board.Move(m, ai.opponentID)
		m_ai := ai.EvaluateMove()
		board.Move(m_ai, ai.myID)
		break
	}

}
