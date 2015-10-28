package main

import (
	"fmt"
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
	Next            *Moves
}

func (m *Moves) Add(row int, startPos int, count int) {
	//fmt.Printf("%d %d %d (%d %d)\n", row, startPos, count, m.NoOfMoves, len(m.AllMoves))
	m.AllMoves[m.NoOfMoves].Row = row
	m.AllMoves[m.NoOfMoves].StartPos = startPos
	m.AllMoves[m.NoOfMoves].Count = count
	m.NoOfMoves++
}

func (m *Moves) AddMoves(row int, col int, moves *Moves) {
	for i := 0; i < moves.NoOfMoves; i++ {
		m.Add(row, moves.AllMoves[i].StartPos+col, moves.AllMoves[i].Count)
	}
}

func (m *Moves) Reset() {
	m.NoOfMoves = 0
}

func (m *Moves) Print() {
	for i := 0; i < m.NoOfMoves; i++ {
		fmt.Printf("%d %d %d (%d)\n", m.AllMoves[i].Row, m.AllMoves[i].StartPos, m.AllMoves[i].Count, m.NoOfMoves)
	}
	fmt.Printf("Done\n")
}
