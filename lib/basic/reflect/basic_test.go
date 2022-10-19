package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type MyInt int

func TestNone(t *testing.T) {
	var b MyInt = 100
	v := reflect.ValueOf(b)
	b2 := v.Interface().(MyInt)
	fmt.Println(b2)
}

func TestTypeValueOf(t *testing.T) {
	var x float64 = 3.1415
	//fmt.Println(reflect.TypeOf(x).String())
	//fmt.Println(reflect.ValueOf(x).String())
	v := reflect.ValueOf(&x)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	fmt.Println("value is can settable : ", v.CanSet())

	vE := v.Elem()
	fmt.Println(vE.Type())
	fmt.Println(vE.Kind())
	fmt.Println("value Ele is can settable : ", vE.CanSet())
	vE.SetFloat(1.23)

	fmt.Println(*v.Interface().(*float64))
	fmt.Println(vE.Interface())
}

type Person struct {
	Name string
	Age  int
}

func TestPersonReflect(t *testing.T) {
	p := &Person{"yaoqijun", 30}
	s := reflect.ValueOf(p).Elem()
	tp := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, tp.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetString("zhao")
	s.Field(1).SetInt(18)
	fmt.Println(p)
}
