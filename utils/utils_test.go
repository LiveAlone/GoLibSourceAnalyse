package main

import (
	"fmt"
	"testing"
	"time"
)

func TestNone(t *testing.T) {
	text := "20230901"
	nt, err := time.Parse("20060102", text)
	fmt.Println(nt.Unix(), err)
}
