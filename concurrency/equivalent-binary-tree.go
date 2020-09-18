package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	tree1 := tree.New(1)

	ch := make(chan int, 10)
	go Walk(tree1, ch)

	fmt.Println(tree1)
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
}
