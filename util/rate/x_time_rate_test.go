package rate

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestRate(t *testing.T) {
	fmt.Println("test time limit rate")

	// 统计rate 限制
	//every := rate.Every(200)
	//l := rate.NewLimiter(every, 1)
	l := rate.NewLimiter(5, 1)
	c, _ := context.WithCancel(context.TODO())
	fmt.Println(l.Limit(), l.Burst())
	for {
		fmt.Println("--------------------------")
		//fmt.Println(l.Allow())
		//rv := l.Reserve()
		//fmt.Println(rv.OK())
		//fmt.Println(rv.Delay())
		err := l.Wait(c)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(10 * time.Millisecond) // 100 次per seconds
		fmt.Println(time.Now().Format("2021-01-02 15:04:05.000"))
	}
}
