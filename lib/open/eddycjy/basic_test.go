package eddycjy

import (
	"fmt"
	"testing"
)

func TestStringEmpty(t *testing.T) {
	var s string
	fmt.Println(s == "")
	fmt.Println(len(s) == 0)
}