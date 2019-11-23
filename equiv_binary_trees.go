package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func WalkRecurse(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	WalkRecurse(t.Left, ch)
	ch <- t.Value
	WalkRecurse(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRecurse(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	res := Same(tree.New(3), tree.New(1))
	fmt.Println(res)
}
