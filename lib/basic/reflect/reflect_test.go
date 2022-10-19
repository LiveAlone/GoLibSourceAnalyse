package reflect

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestImplementsPoint(t *testing.T) {
	wt := reflect.TypeOf((*io.Writer)(nil))
	writerType := wt.Elem()
	ft := reflect.TypeOf((*os.File)(nil))
	//fileType := ft.Elem()
	fmt.Println(ft.Implements(writerType))
}

func printInterfaceInfo (o interface{}) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	fmt.Println(t.Name())
	fmt.Println(v.Kind())
}


func TestValueTypeOf(t *testing.T) {
	a := "hello world"
	printInterfaceInfo(a)
	printInterfaceInfo(&a)
	//printInterfaceInfo(123)
}

func TestFloatSetting(t *testing.T) {
	var x float64 = 3.4
	//v := reflect.ValueOf(x)
	//v.SetFloat(7.1) // Error: will panic.
	//fmt.Println("settability of v:", v.CanSet())

	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	p.Elem().SetFloat(123.66)
	fmt.Println(p)
	fmt.Println(x)
}

func TestTagContentInfo(t *testing.T) {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

	v := reflect.ValueOf(s)
	fmt.Println(v)
	fmt.Print(v.CanSet())
}

func TestStructFieldUpdate(t *testing.T) {
	type T struct {
		A int
		B string
	}

	t1 := T{23, "skidoo"}
	// 通过指针调用， 防止函数参数 内存拷贝, Elem 获取, 指针Value，转换当前对象Value
	s := reflect.ValueOf(&t1).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t1)
}

func TestFunctionMakeCall(t *testing.T) {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwap expects fptr to be a pointer to a nil function.
	// It sets that pointer to a new function created with MakeFunc.
	// When the function is invoked, reflect turns the arguments
	// into Values, calls swap, and then turns swap's result slice
	// into the values returned by the new function.
	makeSwap := func(fptr interface{}) {
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
		fn := reflect.ValueOf(fptr).Elem()

		// Make a function of the right type.
		v := reflect.MakeFunc(fn.Type(), swap)

		// Assign it to the value fn represents.
		fn.Set(v)
	}

	// Make and call a swap function for ints.
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))
}