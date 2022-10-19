package basic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestFunction(t *testing.T) {
	// 函数支持闭包方式, 绑定运行变量
	//pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i),
	//	)
	//}
}

func TestPoint(t *testing.T) {

	//i, j := 42, 2701
	//
	//p := &i         // 指向 i
	//fmt.Println(*p) // 通过指针读取 i 的值
	//*p = 21         // 通过指针设置 i 的值
	//fmt.Println(i)  // 查看 i 的值
	//
	//p = &j         // 指向 j
	//*p = *p / 37   // 通过指针对 j 进行除法运算
	//fmt.Println(j) // 查看 j 的值

	// var
	//v1 := Vertex{1, 2} // 创建一个 Vertex 类型的结构体
	//v2 := Vertex{X: 1} // Y:0 被隐式地赋予
	//v3 := Vertex{}     // X:0 Y:0
	//p := &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	//fmt.Println(v1, v2, v3, p)

	//var s []int // nil 未初始化
	//fmt.Println(s, len(s), cap(s))
	//if s == nil {
	//	fmt.Println("nil!") // print nil
	//}

	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

//type Vertex struct {
//	X, Y int
//}

func TestProcess(t *testing.T) {

	//sum := 0
	//for i := 0; i < 100; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//a := 10
	//if c := a * a; c < 120 {
	//	fmt.Println("test")
	//}

	// switch
	//fmt.Println("Go runs on")
	//switch os := runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	// freebsd, openbsd,
	//	// plan9, windows...
	//	fmt.Printf("%s.\n", os)
	//}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func TestVar(t *testing.T) {

	// 返回相同， seed 需要手动设置
	fmt.Println("import random package value is ", rand.Intn(100))

	fmt.Println("Add value is ", Add(10, 12))
}

func Add(x, y int) int {
	return x + y
}
