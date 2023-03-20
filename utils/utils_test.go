package main

import (
	"fmt"
	"testing"
)

func TestNone(t *testing.T) {
	arrList := []int{501110, 501111, 501112, 501113, 501116, 501117, 501119, 501121, 501127, 501128}
	for i := len(arrList) - 1; i >= 0; i-- {
		fmt.Println(arrList[i])
	}
}
