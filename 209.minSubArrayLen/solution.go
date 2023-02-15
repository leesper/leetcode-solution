package leetcode209

import "math"

// 已知：正整数target，正整数数组nums
// 未知：整数l
// 条件：
// （1）l是nums中连续子数组sub的长度；
// （2）sub中元素之和>=target
// （3）sub的长度是所有连续子数组中最小的；
// （4）1 <= target <= 10^9
// （5）1 <= nums.length <= 10^5
// （6）1 <= nums[i] <= 10^5
func minSubArrayLen(target int, nums []int) int {
	result, sum := math.MaxInt, 0
	left, right := 0, 0

	for right < len(nums) {
		sum += nums[right]
		right++

		for sum >= target {
			if right-left < result {
				result = right - left
			}
			sum -= nums[left]
			left++
		}
	}

	if result == math.MaxInt {
		return 0
	}

	return result
}
