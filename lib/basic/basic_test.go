package basic

import (
	"fmt"
	"testing"
)

type LocalAdder interface {
	LocalAdd(int32, int32)
}
type LocalAdderFunc func(int32, int32)

func (f LocalAdderFunc) LocalAdd(a, b int32) {
	f(a, b)
}

func add(a, b int32, adder LocalAdder) {
	adder.LocalAdd(a, b)
}

func TestSingleFunction(t *testing.T) {
	af := LocalAdderFunc(func(a, b int32) {
		fmt.Println(a, b, a+b)
	})
	add(1, 2, af)
}

func SinglePrint(param string) string {
	fmt.Println("print single line content ", param)
	return "haha"
}

func TestDeferCommand(t *testing.T) {
	defer SinglePrint(SinglePrint("yaoqijun"))
	fmt.Println("none")
}
