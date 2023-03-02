package main

import (
	"log"
	"testing"
)

type Person struct {
	Name string
}

func TestNone(t *testing.T) {
	//p := &Person{
	//	Name: "123",
	//}
	var p *Person
	log.Fatalf("p is :%v", *p)
}
