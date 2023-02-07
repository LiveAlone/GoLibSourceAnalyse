package leetcode

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	pre := 1
	for index := 1; index < len(s); index++ {
		currentSize := 1
		for i := 1; i <= pre; i++ {
			if s[index-i] == s[index] {
				break
			}
			currentSize += 1
		}
		pre = currentSize
		if currentSize > max {
			max = currentSize
		}
	}
	return max
}
