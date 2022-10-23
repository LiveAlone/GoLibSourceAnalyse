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

// nil 可以调用

type T struct{}

func (t *T) Hello() string {
	//if t == nil {
	//	fmt.Println("脑子进煎鱼了")
	//	return ""
	//}
	return "煎鱼进脑子了"
}

func TestPrintNil(t *testing.T) {
	var et T
	fmt.Println(et.Hello())
}

type Person struct {
	Nane string
	Age  int
}

func TestPersonList(t *testing.T) {
	p := []Person{
		{Nane: "yao", Age: 18},
		{Nane: "qi", Age: 19},
		{Nane: "jun", Age: 20},
	}

	fmt.Println(p)
	for i := 0; i < len(p); i++ {
		p[i].Age = 200
	}
	fmt.Println(p)
}
