package eddycjy

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSingle(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("AllFinish")
}