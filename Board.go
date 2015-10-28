package main

import (
	"fmt"
)

var MaxRow int = 5

type Board struct {
	Box             [][]int
	MaxCut          int
	CurrentCutCount int
	//allocator       Allocator
	allocator CustomAllocator
	moveCache map[int]*Moves
}

func (b *Board) Init() {
	b.Box = make([][]int, MaxRow)
	for i := 0; i < MaxRow; i++ {
		b.Box[i] = make([]int, i+1)
	}
	b.MaxCut = MaxRow * (MaxRow + 1) / 2
	b.allocator.Init(1000)
	b.InitMoveCache()
}

func (b *Board) InitMoveCache() {
	b.moveCache = make(map[int]*Moves)
	for i := 0; i < MaxRow; i++ {
		moves := new(Moves)
		b.GetAllMovesOnRow(i, moves)
		b.moveCache[i+1] = moves
	}
	//fmt.Printf("moveCache=%v\n", b.moveCache)
	//fmt.Printf("moveCache[4]=%v\n", b.moveCache[4])
	//fmt.Printf("moveCache[5]=%v\n", b.moveCache[5])
}

func (b *Board) Deinit() {
	b.Box = nil
}

func (b *Board) Print() {
	for i := 0; i < MaxRow; i++ {
		Col := len(b.Box[i])
		for j := 0; j < Col; j++ {
			fmt.Printf("%d ", b.Box[i][j])
		}
		fmt.Printf("\n")
	}
}

func (b *Board) Move(m Move, token int) {
	c := m.StartPos
	for i := 0; i < m.Count; i++ {
		b.Box[m.Row][c] = token
		b.CurrentCutCount++
		c++
	}
}

func (b *Board) UndoMove(m Move) {
	c := m.StartPos
	for i := 0; i < m.Count; i++ {
		b.Box[m.Row][c] = 0
		b.CurrentCutCount--
		c++
	}
}

func (b *Board) isValidMove(m Move) bool {
	c := m.StartPos
	for i := 0; i < m.Count; i++ {
		if b.Box[m.Row][c] != 0 {
			return false
		}
		c++
	}
	return true
}

func (b *Board) isLoser() bool {
	return (b.CurrentCutCount == b.MaxCut)
}

func (b *Board) GetAllMoves() *Moves {
	moves := b.allocator.Capture()
	count := 0
	flagFirst := true
	row, col := 0, 0
	for r := 0; r < MaxRow; r++ {
		count = 0
		for c := 0; c < len(b.Box[r]); c++ {
			if b.Box[r][c] == 0 {
				if flagFirst {
					row = r
					col = c
					flagFirst = false
				}
				count++
			} else {
				if count > 0 {
					moves.AddMoves(row, col, b.moveCache[count])
				}
				count = 0
				flagFirst = true
			}
		}
		if count > 0 {
			moves.AddMoves(row, col, b.moveCache[count])
		}
		flagFirst = true
	}

	//moves.Print()
	return moves
}

func (b *Board) GetAllMovesOnRow(row int, moves *Moves) {
	//result := list.New()
	flagFirst := true
	cFirst := 0
	count := 0
	//fmt.Printf("Length of Box=%d\n", len(b.Box[row]))
	for sc := 1; sc <= row+1; sc++ {

		for c := 0; c < len(b.Box[row]); c++ {
			//fmt.Printf("1 sc=%d c=%d cFirst=%d count=%d flagFirst=%v val=%d\n", sc, c, cFirst, count, flagFirst, b.Box[row][c])

			if count < sc {
				if b.Box[row][c] == 0 {
					if flagFirst {
						cFirst = c
						flagFirst = false
						count = 0
					}
					count++
				} else {
					flagFirst = true
					count = 0
					cFirst = c
					//fmt.Printf("Skip %d %d\n", c, sc)
				}

			} else {

				if count == sc {
					//fmt.Printf("Added %d %d\n", cFirst, sc)
					/*for x := 0; x < sc; x++ {
						l.PushBack(v)
						v++
					}
					result.PushBack(l)
					*/
					moves.Add(row, cFirst, sc)

					c = cFirst
					flagFirst = true
					count = 0
				}
			}
			//	fmt.Printf("2 sc=%d c=%d cFirst=%d count=%d flagFirst=%v val=%d\n", sc, c, cFirst, count, flagFirst, b.Box[row][c])

		} //end of for(c)
		if count == sc {
			//	fmt.Printf("Added %d %d\n", cFirst, sc)
			/*for x := 0; x < sc; x++ {
				l.PushBack(v)
				v++
			}
			result.PushBack(l)
			*/
			moves.Add(row, cFirst, sc)

			flagFirst = true
			count = 0
		}

		count = 0
		flagFirst = true
	}
	//fmt.Println(result)
}
