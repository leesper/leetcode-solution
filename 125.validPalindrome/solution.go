package leetcode125

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = strings.ToLower(reg.ReplaceAllString(s, ""))
	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
