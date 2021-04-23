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

	if t.Right != nil {
		Walk(t.Right, ch)
	}

	ch <-t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)

	ch := make(chan int)
	go Walk(t1, ch)

	var same bool
	for i := 0; i < 10; i++ {
		same = binarySearch(<-ch, t2)
	}

	return same
}

func binarySearch(n int, t *tree.Tree) bool {
	if t.Value == n {
		return true
	}

	var same bool
	if t.Left != nil {
		same = binarySearch(n, t.Left)
	}

	if t.Right != nil {
		same = binarySearch(n, t.Right)
	}

	return same
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
