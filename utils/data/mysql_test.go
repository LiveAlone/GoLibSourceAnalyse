package data

import (
	"fmt"
	"testing"
)

func TestQueryColumns(t *testing.T) {
	url := "homework:homework@tcp(10.112.36.52:6060)/information_schema?charset=utf8mb4&parseTime=True&loc=Local"
	columns, err := QueryColumns(url, "hxx_apps", "tblActivityBasic")
	if err != nil {
		t.Error(err)
	}
	for i := range columns {
		fmt.Println(i, columns[i])
	}
}
