package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	v := x
	rs := 0
	for v > 0 {
		rs = rs*10 + (v % 10)
		v = v / 10
	}
	return rs == x
}

func myAtoi(s string) int {
	curs := strings.TrimSpace(s)
	rs, pos := 0, 0
	neg := false
	if curs[0] == '-' {
		pos += 1
		neg = true
	}
	for pos < len(curs) && curs[pos] >= '0' && curs[pos] <= '9' {
		if rs*10 < rs {
			break
		}
		rs = rs*10 + int(curs[pos]-'0')
		pos += 1
	}
	if neg {
		return -rs
	}
	return rs
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
