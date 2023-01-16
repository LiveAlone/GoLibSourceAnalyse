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
	var body Body
	s := "{\"yao\":123,\"qi\":456,\"jun\":[1,2,3]}"
	err := json.Unmarshal([]byte(s), &body.Msg)
	fmt.Println(err)
	fmt.Println(body)
}
