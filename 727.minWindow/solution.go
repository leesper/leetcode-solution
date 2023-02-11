package leetcode727

import "math"

// 已知：字符串s，t
// 未知：字符串w
// 条件：
// （1）w是s的最短子串
// （2）t是w的子序列
// （3）若s中没有窗口包含t中所有字符，返回空串
// （4）如果有不止一个最短长度窗口，返回最靠左的1个
// （5）1 <= s <= 20000，只包含小写字母
// （6）1 <= t <= 100，只包含小写字母
func minWindow(s string, t string) string {
	result := ""
	min := math.MaxInt
	sIdx := make([]int, len(t))
	tSeq := make([]rune, len(t))

	for i, c := range t {
		tSeq[i] = c
		sIdx[i] = -1
	}

	left, right := 0, 0
	prevChar, next := rune(0), 0
	for right < len(s) {
		if rune(s[right]) == tSeq[next] {
			sIdx[next] = right
			prevChar = tSeq[next]
			next++
		} else if rune(s[right]) == prevChar {
			sIdx[next-1] = right
		}
		right++

		for next >= len(tSeq) {
			if len(s[sIdx[0]:right]) < min {
				result = s[sIdx[0]:right]
				min = len(s[sIdx[0]:right])
			}
			updated := 0
			left = sIdx[0] + 1
			prevChar = rune(0)
			for left < right {
				if rune(s[left]) == tSeq[updated] {
					sIdx[updated] = left
					prevChar = tSeq[updated]
					updated++
				} else if rune(s[left]) == prevChar {
					sIdx[updated-1] = left
				}
				left++
			}
			next = updated
		}
	}
	return result
}
