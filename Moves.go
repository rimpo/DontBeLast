package main

type Move struct {
	Row      int
	StartPos int
	Count    int
}

type Moves struct {
	AllMoves        [40]Move
	NoOfMoves       int
	NoOfMovesPlayed int
}

func (m *Moves) Add(row int, startPos int, count int) {
	m.AllMoves[m.NoOfMoves].Row = row
	m.AllMoves[m.NoOfMoves].StartPos = startPos
	m.AllMoves[m.NoOfMoves].Count = count
	m.NoOfMoves++
}

func (m *Moves) Reset() {
	m.NoOfMoves = 0
}
