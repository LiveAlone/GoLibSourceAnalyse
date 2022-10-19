package basic

import (
	"fmt"
	"math"
	"testing"
)

func TestAssert(t *testing.T) {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TestNil(t *testing.T) {
	var i I

	var nt *T
	i = nt
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

type Vertex struct {
	X, Y float64
	Z    float64
}

func (v Vertex) Abs() float64 {
	v.Z = math.Sqrt(v.X*v.X + v.Y*v.Y)
	return v.Z
}

func (v *Vertex) AbsPoint() float64 {
	v.Z = math.Sqrt(v.X*v.X + v.Y*v.Y)
	return v.Z
}

func TestStruct(t *testing.T) {
	v := &Vertex{3, 4, 0}
	//fmt.Println(v.Abs())
	fmt.Println(v.AbsPoint())
	fmt.Println(v)
}
