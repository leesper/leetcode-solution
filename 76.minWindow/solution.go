package leetcode76

import (
	"math"
)

// 已知：字符串s和字符串t
// 未知：字符串q
// 条件：
// * q是s的子串
// * q包含t中所有字符且长度最小
// * 对于t中重复字符，q中该字符数量必须不少于t中该字符数量
// * 若s中存在这样的q，则q唯一；若不存在，则q为空串
// * s和t的长度都在1到10^5之间（included）
// * s和t由英文字母组成
// * 算法复杂度要求O(m+n)，m和n分别为字符串s和t的长度
func minWindow(s string, t string) string {
	minStr := ""
	minLen := math.MaxInt

	tCount := make(map[rune]int, len(t))
	for _, c := range t {
		tCount[c]++
	}

	left, right := 0, 0
	window := make(map[rune]int, len(t))
	enough := make(map[rune]bool, len(t))
	for right < len(s) {
		if tCount[rune(s[right])] > 0 {
			window[rune(s[right])]++
			if window[rune(s[right])] >= tCount[rune(s[right])] {
				enough[rune(s[right])] = true
			}
		}
		right++

		for len(window) > 0 {
			if window[rune(s[left])] > 0 {
				if window[rune(s[left])] <= tCount[rune(s[left])] {
					break
				} else {
					window[rune(s[left])]--
				}
			}
			left++
		}

		if len(enough) == len(tCount) {
			str := s[left:right]
			if len(str) <= minLen {
				minLen = len(str)
				minStr = str
			}
		}
	}

	return minStr
}
