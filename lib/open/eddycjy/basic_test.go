package eddycjy

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringEmpty(t *testing.T) {
	var s string
	fmt.Println(s == "")
	fmt.Println(len(s) == 0)
}

// 1.18 提供Cut
func TestCut(t *testing.T) {
	fmt.Println(strings.Cut("yaoqijunmail@foxmail", "@"))
}