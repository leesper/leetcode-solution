package leetcode239

import (
	"container/list"
)

// 输入：整数数组nums，整数k
// 输出：整数数组maxes
// 条件：
// （1）滑动窗口每次向右移动1位
// （2）maxes是nums中大小为k的滑动窗口中最大值集合
// （3）1 <= len(nums) <= 10^5
// （4）-10^4 <= nums[i] <= 10^4
// （5）1 <= k <= len(nums)
func maxSlidingWindow(nums []int, k int) []int {
	maxes := make([]int, len(nums)-k+1)
	deque := list.New()

	left, right := 0, 0
	for right < len(nums) {
		for 0 < deque.Len() && deque.Back().Value.(int) < nums[right] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(nums[right])
		right++

		for right-left >= k {
			maxes[left] = deque.Front().Value.(int)
			if nums[left] == deque.Front().Value.(int) {
				deque.Remove(deque.Front())
			}
			left++
		}
	}
	return maxes
}
