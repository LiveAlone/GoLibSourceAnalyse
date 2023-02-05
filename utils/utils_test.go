package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Body struct {
	Msg interface{}
}

func TestNone(t *testing.T) {
	var entity map[string][]int
	s := "{\"qi\":[10, 100],\"jun\":[1,2,3]}"
	err := json.Unmarshal([]byte(s), entity)
	fmt.Println(err)
	fmt.Println(entity)
}
