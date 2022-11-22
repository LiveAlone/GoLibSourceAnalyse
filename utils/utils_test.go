package main

import (
	"fmt"
	"testing"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/cmd"
)

func TestDemo(t *testing.T) {
	rs := cmd.GenerateFromTable("", "")
	fmt.Println(rs)
}
