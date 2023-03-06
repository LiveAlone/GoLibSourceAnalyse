package main

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type Person struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestNone(t *testing.T) {
	p := &Person{
		Id:   1,
		Name: "yqj",
		Age:  10,
	}
	pj, err := jsoniter.Marshal(p)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(pj))
}
