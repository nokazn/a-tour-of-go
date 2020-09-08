package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	n := 1
	testTree := tree.New(n)
	ch := make(chan int)
	go Walk(testTree, ch)

	nums := make(map[int]bool)
	// ch 内の数値を連想配列に格納
	for i := range ch {
		nums[i] = true
	}

	const K = 10
	ok := true
	// 1 ~ n*K までの数字がすべて含まれているか
	for i := 1; i <= n*K; i++ {
		ok = nums[i] && ok
	}

	if !ok {
		t.Errorf("error!")
	}
}
