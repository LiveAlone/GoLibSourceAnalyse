package main

import (
	"fmt"
	"testing"
)

func TestNone(t *testing.T) {
	a := 1100
	switch a {
	case 100:
		fmt.Println(100)
	case 200:
		fmt.Println(200)
	case 300:
		fmt.Println(300)
	case 400:
		fmt.Println(400)
	default:
		fmt.Println("err")
	}
}
