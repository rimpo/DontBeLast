package main

import (
	"container/list"
)

type Allocator struct {
	freeList list.List
}

func (a *Allocator) Init(n int) {
	for i := 0; i < n; i++ {
		moves := new(Moves)
		a.freeList.PushBack(moves)
	}
}

func (a *Allocator) Capture() *Moves {
	val := a.freeList.Front()
	if val == nil {
		return new(Moves)
	}
	a.freeList.Remove(val)
	return val.Value.(*Moves)
}

func (a *Allocator) Release(moves *Moves) {
	moves.Reset()
	a.freeList.PushBack(moves)
}
