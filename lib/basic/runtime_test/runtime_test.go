package runtime_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestBasic(t *testing.T) {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}