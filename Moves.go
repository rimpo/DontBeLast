package main

import (
//	"fmt"
)

type Move struct {
	Row      int
	StartPos int
	Count    int
}

type Moves struct {
	AllMoves        [64]Move
	NoOfMoves       int
	NoOfMovesPlayed int
}

func (m *Moves) Add(row int, startPos int, count int) {
	//fmt.Printf("%d %d %d (%d %d)\n", row, startPos, count, m.NoOfMoves, len(m.AllMoves))
	m.AllMoves[m.NoOfMoves].Row = row
	m.AllMoves[m.NoOfMoves].StartPos = startPos
	m.AllMoves[m.NoOfMoves].Count = count
	m.NoOfMoves++
}

func (m *Moves) Reset() {
	m.NoOfMoves = 0
}
