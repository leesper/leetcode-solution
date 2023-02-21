package leetcode1004

// 已知：二进制数组nums，整数k
// 未知：1个整数
// 条件：要求的整数为nums数组中翻转k个0后连续1的最大个数
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	max, timesOfZero := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			timesOfZero++
		}

		right++

		if timesOfZero <= k && max < right-left {
			max = right - left
		}

		for timesOfZero > k {
			if nums[left] == 0 {
				timesOfZero--
			}
			left++
		}
	}

	return max
}
