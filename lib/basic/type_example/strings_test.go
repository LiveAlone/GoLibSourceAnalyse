package type_example

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStringChar(t *testing.T) {

	//str := "魑魅魍魉"
	//fmt.Println(len(str))
	//
	//for i, i2 := range str{
	//	fmt.Printf("%d, %#U \n", i, i2)
	//	fmt.Println(strconv.Itoa(int(i2)))
	//}
	//
	//strings.Contains(str, "姚")

	s := strconv.QuoteRune('☺')
	fmt.Println(s)
}