package sync_example

import (
	"fmt"
	"sync"
	"testing"
)

// test sync.Pool borrow object
func TestSyncPool(t *testing.T) {
	type Person struct {
		Name string
	}

	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return &Person{Name: "Name"}
		},
	}

	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)

	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	pool.Put(p)

	// Cache 缓存方式
	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*Person))
}