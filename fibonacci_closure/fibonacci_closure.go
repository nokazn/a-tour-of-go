package main

import "fmt"

func fibonacci() func() int {
	prev := 0
	curr := 1
	return func() int {
		result := curr + prev
		prev = curr
		curr = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
