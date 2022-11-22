package data

import (
	"fmt"
	"testing"

	"github.com/gobeam/stringy"
)

func TestCaseConvert(t *testing.T) {
	s := "yaoQiJun"
	sy := stringy.New(s)
	fmt.Println(sy.SnakeCase().UcFirst())
	fmt.Println(sy.KebabCase().LcFirst())
}
