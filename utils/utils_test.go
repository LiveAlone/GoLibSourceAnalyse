package main

import (
	"fmt"
	"testing"
)

func TestNone(t *testing.T) {
	PrintLocal()
}

func PrintLocal() (rs int) {
	defer func() {
		fmt.Println(rs)
	}()
	return 100
}
