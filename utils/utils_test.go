package main

import (
	"fmt"
	"testing"
)

func TestNone(t *testing.T) {
	n := new(None)
	fmt.Println(TwoValue(1, 2, n.ADD))
}

type None struct {
}

func (n *None) ADD(a, b int) int {
	return 100
}

func TwoValue(a, b int, f func(a, b int) int) int {
	return f(a, b)
}
