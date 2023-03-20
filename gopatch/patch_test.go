package gopatch

import (
	"errors"
	"fmt"
	"testing"
)

func TestNone(t *testing.T) {
	err := errors.New(fmt.Sprintf("hello"))
	// err := fmt.Errorf("hello")
	fmt.Println(err)
}

func TestT2(t *testing.T) {
	foo(3)
	var x = 100
	foo(x)
	foo(1 + 2 + 3 + 4 + 5)
}

func foo(v int) int {
	return v
}

func bar(v int, exist bool) int {
	return v
}
