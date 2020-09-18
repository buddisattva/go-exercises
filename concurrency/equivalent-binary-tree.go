package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

const capacity = 10

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
	ch1 := make(chan int, capacity)
	go Walk(t1, ch1)

	ch2 := make(chan int, capacity)
	go Walk(t2, ch2)

	for i := 0; i < capacity; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	tree1 := tree.New(1)
	traverse(tree1)
	tree2 := tree.New(1)
	traverse(tree2)

	result := Same(tree1, tree2)
	fmt.Println("---")
	fmt.Printf("Are above trees same as each other? %v\n", result)

	tree2 = tree.New(2)

	result = Same(tree1, tree2)
	fmt.Println("---")
	fmt.Println("Tree: ", tree1)
	fmt.Println("Tree: ", tree2)
	fmt.Printf("Are above trees same as each other? %v\n", result)
}

func traverse(tree *tree.Tree) {
	ch := make(chan int, capacity)
	go Walk(tree, ch)

	fmt.Println("Tree: ", tree)
	fmt.Println("Traversal:")
	for i := 0; i < capacity; i++ {
		fmt.Println(<-ch)
	}
}
