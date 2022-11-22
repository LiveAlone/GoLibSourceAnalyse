package main

import (
	"fmt"
	"testing"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/cmd"
)

func TestDemo(t *testing.T) {
	rs := cmd.GenerateFromTable("hxx_live", "tblTeacherLiveQuestionInfo")
	fmt.Println(rs)
}
