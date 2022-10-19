package rate

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/ratelimit"
)

func TestRateLimit(t *testing.T) {
	rl := ratelimit.New(10) // per second

	prev := time.Now()
	for i := 0; i < 100; i++ {
		now := rl.Take()
		if i > 0 {
			fmt.Println(i, now.Sub(prev))
		}
		prev = now
	}
}
