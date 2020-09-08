package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

const LENGTH = 10

func walker(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walker(t.Left, ch)
	walker(t.Right, ch)
	ch <- t.Value
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walker(t, ch)
	close(ch)
}

func makeMap(ch chan int) map[int]bool {
	nums := make(map[int]bool)
	for i := range ch {
		nums[i] = true
	}
	return nums
}

func makeSlice(ch chan int) []int {
	slice := []int{}
	for i := range ch {
		slice = append(slice, i)
	}
	return slice
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	t1Map := makeMap(ch1)
	t2Slice := makeSlice(ch2)

	ok := true
	for _, v := range t2Slice {
		ok = ok && t1Map[v]

		if !ok {
			break
		}
	}

	return ok
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(3)))
}
