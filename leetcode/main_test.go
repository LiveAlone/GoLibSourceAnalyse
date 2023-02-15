package leetcode

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	fmt.Println(convert("A", 1))
}

func convert(s string, numRows int) string {
	rs := make([]byte, 0, len(s))
	step := 2*numRows - 2
	if step == 0 {
		return s
	}
	for pos := 0; pos <= step>>1; pos++ {
		if pos == 0 || pos == step>>1 {
			for from := pos; from < len(s); from += step {
				rs = append(rs, s[from])
			}
		} else {
			for from := pos; from < len(s); from += step {
				rs = append(rs, s[from])
				nextPos := from - pos*2 + step
				if nextPos < len(s) {
					rs = append(rs, s[nextPos])
				}
			}
		}
	}
	return string(rs)
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
