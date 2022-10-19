package sync_example

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestConvertPointerType(t *testing.T) {
	var a float64
	a = 100.123
	// 强制转换
	pointer := *(*uint64)(unsafe.Pointer(&a))
	fmt.Println(pointer)
	fmt.Println(a)
}

func TestReflect(t *testing.T) {
	u := reflect.ValueOf(new(int)).Pointer()
	p := (*int)(unsafe.Pointer(u))
	fmt.Println(*p)
}

func TestBasic(t *testing.T) {

	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}

	p := unsafe.Pointer(&s)
	up0 := uintptr(p)
	pb := (*byte)(p)
	*pb = 10
	fmt.Println(s)
	fmt.Println(p, up0, pb)

	up := up0 + unsafe.Offsetof(s.b)
	p = unsafe.Pointer(up)
	pb = (*byte)(p)
	*pb = 20
	fmt.Println(s)
	fmt.Println(p, up, pb)

	up = up0 + unsafe.Offsetof(s.c)
	p = unsafe.Pointer(up)
	pb = (*byte)(p)
	*pb = 30
	fmt.Println(s)

	up = up0 + unsafe.Offsetof(s.d)
	p = unsafe.Pointer(up)
	pi := (*int64)(p)
	*pi = 40
	fmt.Println(s)
}

func TestStringReader(t *testing.T) {
	sr := strings.NewReader("abcdef")
	fmt.Println(sr)
	p := unsafe.Pointer(sr)
	up0 := uintptr(p)
	if sf, ok := reflect.TypeOf(*sr).FieldByName("i"); ok {
		up := up0 + sf.Offset
		p = unsafe.Pointer(up)
		pi := (*int64)(p)
		*pi = 3 // 修改索引, 指定当前读取数值
	}
	// 看看修改结果
	fmt.Println(sr)
	b, err := sr.ReadByte()
	fmt.Printf("%c, %v\n", b, err)
}

type LocalReader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func TestReaderUpdate(t *testing.T) {
	sr := strings.NewReader("abcdef")
	fmt.Println(sr)
	p := unsafe.Pointer(sr)
	pR := (*LocalReader)(p)
	// 这样就可以自由修改 sr 中的私有成员了, 类似地址强制修改变量类型
	(*pR).i = 3 // 修改索引
	fmt.Println(sr)
	b, err := sr.ReadByte()
	fmt.Printf("%c, %v\n", b, err)
}
