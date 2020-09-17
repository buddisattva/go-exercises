package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	beforePrev := 0
	prev := 0
	cur := 0

	return func() int {
		if cur == 0 {
			cur = 1
			return 0
		} else if cur == 1 && prev == 0 {
			cur = 1
			prev = 1
			return cur
		}

		cur = prev + beforePrev
		beforePrev = prev
		prev = cur
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
