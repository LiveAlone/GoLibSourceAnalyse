package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Local struct {
	Names []string
}

func TestNone(t *testing.T) {
	l := Local{}
	if l.Names == nil {
		fmt.Println("none")
	}
	rs, _ := json.Marshal(l.Names)
	fmt.Println(string(rs))
}
