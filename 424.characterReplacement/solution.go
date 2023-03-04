package leetcode424

// 未知量：子字符串长度l
// 已知数据：字符串s和整数k
// 条件：
// （1）l为在s中最多执行k次替换后包含相同字母的最长子字符串长度
// （2）s长度在1到1e5之间
// （3）s由大写英文字母构成
// （4）k的长度在0到len(s)之间
func characterReplacement(s string, k int) int {
	left, right, max := 0, 0, 0
	kindsOfLetters := map[byte]int{}

	for right < len(s) {
		kindsOfLetters[s[right]]++
		right++

		if minimumFlips(kindsOfLetters) <= k && max < right-left {
			max = right - left
		}

		for minimumFlips(kindsOfLetters) > k {
			kindsOfLetters[s[left]]--
			if kindsOfLetters[s[left]] == 0 {
				delete(kindsOfLetters, s[left])
			}
			left++
		}
	}

	return max
}

func minimumFlips(count map[byte]int) int {
	if len(count) <= 1 {
		return 0
	}

	sum := 0
	max := 0
	for _, cnt := range count {
		if cnt > max {
			max = cnt
		}
		sum += cnt
	}

	return sum - max
}
