package leetcode1

import "sort"

// 已知：整数数组nums，整数target
// 未知：下标i，j
// 条件：
// （1）nums[i] + nums[j] == target
// （2）每种输入只会对应1个答案
// （3）数组中同一个元素不会在答案中重复出现
// （4）要求时间复杂度小于O(n^2)
func twoSum(nums []int, target int) []int {
	indexes := map[int][]int{}

	for i, v := range nums {
		if _, ok := indexes[v]; !ok {
			indexes[v] = []int{i}
		} else {
			indexes[v] = append(indexes[v], i)
		}
	}

	sort.Ints(nums)
	firstIndex, secondIndex := -1, -1
	for i := range nums {
		j := binaryFind(nums, i+1, target-nums[i])

		if j != -1 {
			firstIndex = indexes[nums[i]][0]
			indexes[nums[i]] = indexes[nums[i]][1:]
			secondIndex = indexes[nums[j]][0]
			if secondIndex < firstIndex {
				firstIndex, secondIndex = secondIndex, firstIndex
			}
			break
		}
	}

	return []int{firstIndex, secondIndex}
}

func binaryFind(nums []int, start, target int) int {
	left, right := start, len(nums)-1

	middle := -1
	for left <= right {
		middle = (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
