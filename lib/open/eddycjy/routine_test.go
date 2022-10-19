package eddycjy

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestRoutineCounter(t *testing.T) {
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				// count 并发不可见  go race 分析
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func TestGoRoutine(t *testing.T) {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	// chan 需要初始化， 否则一直阻塞
	//var ch chan int
	ch := make(chan int)
	go func() {
		c := <-ch
		fmt.Println(c)
	}()

	ch <- 100
	time.Sleep(time.Second)
	fmt.Println("finish")
}