package main

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
	Age  int32
}

type Person struct {
	Stu Student
}

func TestPool(t *testing.T) {
	p := &Person{}
	fmt.Println(p)
}
