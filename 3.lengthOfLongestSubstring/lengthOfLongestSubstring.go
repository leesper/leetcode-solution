package leetcode3

// 执行时间0ms，击败100%用户，空间2.3M，击败95%用户
// 小tip：英文数字+字母+符号+空格是在暗示你这是ascii码，所以用一个长度为128的数组就可以记录所有字符的位置
// 不要使用make或者new这样的内置函数动态创建数组，尽量编译时定下来，就可以减少时间上的损耗
// 不追求代码特别短，只追求可读性好
func lengthOfLongestSubstring(s string) int {
	const ASCII = 128
	lenOfLongestSubstr := 0
	var pos [ASCII]int
	for b := 0; b < len(pos); b++ {
		pos[b] = -1
	}

	max := func(m, n int) int {
		if m > n {
			return m
		}
		return n
	}

	i := 0
	for j := 0; j < len(s); j++ {
		if pos[s[j]] >= 0 {
			i = max(i, pos[s[j]]+1)
		}
		pos[s[j]] = j
		lenOfLongestSubstr = max(lenOfLongestSubstr, j-i+1)
	}

	return lenOfLongestSubstr
}
