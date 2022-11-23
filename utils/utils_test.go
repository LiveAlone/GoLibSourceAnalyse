package main

import (
	"fmt"
	"testing"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/cmd"
)

func TestDemo(t *testing.T) {
	url := "homework:homework@tcp(10.112.36.52:6060)/information_schema?charset=utf8mb4&parseTime=True&loc=Local"
	rs := cmd.GenerateFromTable(url, "hxx_live", "tblTeacherLiveQuestionInfo")
	fmt.Println(rs)
}
