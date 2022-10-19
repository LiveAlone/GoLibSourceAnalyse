package eddycjy

import (
	"fmt"
	"reflect"
	"testing"
)

type People struct {}

type MyStruct struct {
	Name string
}

func (s MyStruct) SetName1(name string) {
	s.Name = name
}

func (s *MyStruct) SetName2(name string) {
	s.Name = name
}

func TestInterfaceNil(t *testing.T) {
	var v interface{}
	var point *int
	fmt.Println(point, point == nil)
	fmt.Println(v, v == nil)
	v = point
	fmt.Println(v, v == nil)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func TestStructCompare(t *testing.T) {

	////////// 1

	//a := &People{}
	//b := &People{}
	//println(a, b, a == b)
	//
	//c := &People{}
	//d := &People{}
	//fmt.Println(c, d)
	//println(c, d, c == d)

	// * -gcflags="-N -l"
	// * People 新增 Name string 字段

	////////////// 2 不同使用方式
	m := MyStruct{"yaoqijun"}
	fmt.Println(m)
	m.SetName1("name1")
	fmt.Println(m)	// (s MyStruct, name string)
	m.SetName2("name2")
	fmt.Println(m)	// (s *MyStruct, name string)
}
