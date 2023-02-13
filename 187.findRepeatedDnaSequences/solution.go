package leetcode187

// 已知：字符串s，整数k
// 未知：字符串的集合S
// 条件：
// （1）S由s的子串构成
// （2）S中的元素长度都为k
// （3）S中的元素在s中出现超过1次
// （4）0 <= s.length <= 10^5
// （5）s表示DNA序列，即s[i]=='A'、'C'、'G' or 'T'
func findRepeatedDnaSequences(s string, k int) []string {
	subseqCount := map[string]int{}
	results := []string{}

	left, right := 0, 0
	for right < len(s) {
		right++

		for right-left >= k {
			subseqCount[s[left:right]]++
			left++
		}
	}

	for str, cnt := range subseqCount {
		if cnt > 1 {
			results = append(results, str)
		}
	}
	return results
}
