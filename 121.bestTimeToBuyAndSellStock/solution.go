package leetcode121

// 已知：表示每日股票价格的整数数组prices
// 未知：表示利润的整数maxProfit
// 条件：
// （1）只能选择某天买入，并在未来的某天卖出
// （2）maxProfit是买入卖出后能获取的最大利润
// （3）不能获取任何利润，则返回0
// （4）1 <= prices.length <= 10^5
// （5）0 <= prices[i] <= 10^4，表示一支给定股票第i天价格
// 问题转化为：在prices中找到两个索引i和j，i < j，且prices[j] - prices[i]最大，返回最大值
func maxProfit(prices []int) int {
	left, right, result := 0, 0, 0

	for right < len(prices) {
		if prices[right]-prices[left] > result {
			result = prices[right] - prices[left]
		}

		for prices[left] > prices[right] {
			left++
		}

		right++
	}
	return result
}
