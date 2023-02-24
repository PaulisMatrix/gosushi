package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t1 *Tree) New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

//issa binary tree
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}

	return true
}

func GoWalk() {
	ch := make(chan int)
	var t *Tree

	go Walk(t.New(1), ch)
	for val := range ch {
		fmt.Print(val)
	}

	if Same(t.New(1), t.New(1)) {
		fmt.Println("OK")
	}

	if !Same(t.New(1), t.New(2)) {
		fmt.Println("OK")
	}

}
