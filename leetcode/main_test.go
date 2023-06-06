package leetcode

import (
	"testing"
)

func TestBasic(t *testing.T) {
}

func coinChange(coins []int, amount int) int {
	if amount == 0 || len(coins) == 0 {
		return 0
	}
	rs := coinChange(coins[1:len(coins)], amount)
	
	return rs
}
