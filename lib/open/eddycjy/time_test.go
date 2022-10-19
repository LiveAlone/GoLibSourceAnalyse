package eddycjy

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeAfter(t *testing.T) {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		ch <- "输出内容"
	}()

	for  {
		select {
		case value := <-ch:
			fmt.Println("gain value is ", value)
		case <-time.After(time.Second * 1):
			fmt.Println("超时了！！！")
		}
	}
}
