package main

import (
	"container/list"
	"fmt"
)

var MaxRow int = 5

type Board struct {
	Box [][]int
}

func (b *Board) Init() {
	b.Box = make([][]int, MaxRow)
	for i := 0; i < MaxRow; i++ {
		b.Box[i] = make([]int, i+1)
	}
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

func (b *Board) Move(row int, startPos int, count int) {
	c := startPos
	for i := 0; i < count; i++ {
		b.Box[row][c] = 1
		c++
	}
}

func (b *Board) GetAllMovesOnRow(row int) *list.List {
	result := list.New()
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
					l := list.New()
					//fmt.Printf("Added %d %d\n", cFirst, sc)
					v := cFirst
					for x := 0; x < sc; x++ {
						l.PushBack(v)
						v++
					}
					result.PushBack(l)

					c = cFirst
					flagFirst = true
					count = 0
				}
			}
			//	fmt.Printf("2 sc=%d c=%d cFirst=%d count=%d flagFirst=%v val=%d\n", sc, c, cFirst, count, flagFirst, b.Box[row][c])

		} //end of for(c)
		if count == sc {
			l := list.New()

			//	fmt.Printf("Added %d %d\n", cFirst, sc)
			v := cFirst
			for x := 0; x < sc; x++ {
				l.PushBack(v)
				v++
			}

			result.PushBack(l)
			flagFirst = true
			count = 0
		}

		count = 0
		flagFirst = true
	}
	//fmt.Println(result)
	return result
}

func (b *Board) PrintAllMoves() {
	for i := 0; i < MaxRow; i++ {
		result := b.GetAllMovesOnRow(i)
		for e := result.Front(); e != nil; e = e.Next() {
			//fmt.Println(e.Value)
			val := e.Value.(*list.List)
			for v := val.Front(); v != nil; v = v.Next() {
				fmt.Printf("%d ", v.Value)
			}
			fmt.Printf("\n")
		}

		fmt.Printf("Total Moves (%d): %d\n", i, result.Len())
	}

}
