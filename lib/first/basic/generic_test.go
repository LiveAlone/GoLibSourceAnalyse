package basic

import (
	"fmt"
	"testing"
)

type LocalNumber interface {
	int64 | float64
}

func GenericAdd2[V LocalNumber](list []V) V {
	var sum V
	for _, num := range list {
		sum += num
	}
	return sum
}

func GenericAdd[V int64 | float64](list []V) V {
	var sum V
	for _, num := range list {
		sum += num
	}
	return sum
}

func TestGenericBasic(t *testing.T) {
	a1 := []int64{1, 2, 3, 4, 5}
	fmt.Println(GenericAdd2(a1))

	b1 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(GenericAdd2(b1))
}
