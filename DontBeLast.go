package main

import (
	//	"container/list"
	"container/list"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var movesMap map[string]*list.List
var MAXROW = 5

func fetchAllMoves(curr_state string, moves *list.List, max_steps int) {
	for i := 1; i <= max_steps; i++ {
		fetchMoves(curr_state, moves, i)
	}
}

func fetchMoves(curr_state string, moves *list.List, steps int) {
	if len(curr_state) <= steps-1 {
		return
	}

	val := []byte(curr_state)

	for i, _ := range val {
		if i+steps-1 < len(val) {
			success := true
			for k := 0; k < steps; k++ {
				if val[i+k] == '0' {
					success = false
				}
			}
			if success {
				for k := 0; k < steps; k++ {
					val[i+k] = '0'
				}

				moves.PushBack(string(val))

				for k := 0; k < steps; k++ {
					val[i+k] = '1'
				}
			}
		}
	}
}

func recursivePopulateMoveMap(curr_state string) {
	l := list.New()
	fetchAllMoves(curr_state, l, MAXROW)
	movesMap[curr_state] = l
	for e := l.Front(); e != nil; e = e.Next() {
		if val, ok := e.Value.(string); ok {
			if _, ok := movesMap[val]; ok {
			} else {
				recursivePopulateMoveMap(val)
			}
		}
	}
}

func createMovesMap() {
	recursivePopulateMoveMap("1")
	recursivePopulateMoveMap("11")
	recursivePopulateMoveMap("111")
	recursivePopulateMoveMap("1111")
	recursivePopulateMoveMap("11111")
}

func displayBoard(board *[]string) {
	for i := 0; i < MAXROW; i++ {
		fmt.Printf("%s\n", (*board)[i])
	}
}

func recursivePlay(curr_move string, curr_player int, board *[]string) int {
	old_state := (*board)[len(curr_move)-1]
	(*board)[len(curr_move)-1] = curr_move

	//fmt.Printf("curr_player=%d curr_move=%s\n", curr_player, curr_move)
	//displayBoard(board)
	next_player := 0
	if curr_player == 1 {
		next_player = 2
	} else {
		next_player = 1
	}

	moveCount := 0
	for i := 0; i < MAXROW; i++ {
		l := movesMap[(*board)[i]]
		moveCount += l.Len()
	}

	if moveCount == 0 {
		(*board)[len(curr_move)-1] = old_state
		if curr_player == 1 {
			//fmt.Printf("Win\n")
			return 1
		} else {
			//fmt.Printf("Lost\n")
			return 0
		}
	}

	totalScore := 0
	for i := 0; i < MAXROW; i++ {
		l := movesMap[(*board)[i]]
		for e := l.Front(); e != nil; e = e.Next() {
			//fmt.Printf("%s ", e.Value)
			if val, ok := e.Value.(string); ok {
				totalScore += recursivePlay(val, next_player, board)
			}
		}
	}
	(*board)[len(curr_move)-1] = old_state
	return totalScore
}
func play(board *[]string) {

	var moveScore map[string]int
	moveScore = make(map[string]int)

	for i := 0; i < MAXROW; i++ {
		l := movesMap[(*board)[i]]
		for e := l.Front(); e != nil; e = e.Next() {
			//fmt.Printf("%s ", e.Value)
			if val, ok := e.Value.(string); ok {
				score := recursivePlay(val, 1, board)
				moveScore[val] = score
				//fmt.Printf("score=%d\n", score)
			}
		}
	}

	for k, v := range moveScore {
		fmt.Printf("\n%s = %d", k, v)
	}
}

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

	movesMap = make(map[string]*list.List)
	createMovesMap()

	board := make([]string, 5)
	board[0] = "1"
	board[1] = "11"
	board[2] = "111"
	board[3] = "1111"
	board[4] = "11111"

	play(&board)

	fmt.Printf("\n ------------\n")

	return

	for k, v := range movesMap {
		fmt.Printf("\n%s = ", k)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Printf("%s ", e.Value)
		}
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
}
