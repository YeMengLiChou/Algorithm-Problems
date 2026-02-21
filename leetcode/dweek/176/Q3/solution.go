package main

func rob(nums []int, colors []int) int64 {
	n := len(nums)
	dp := make([]int64, n+1)
	dp[0] = int64(nums[0])
	for i := 1; i < n; i++ {
		if colors[i] != colors[i-1] { // 都拿
			dp[i] = dp[i-1] + int64(nums[i])
		} else { // 需要从上一个颜色中转移
			if i >= 2 {
				dp[i] = max(dp[i-1], dp[i-2]+int64(nums[i]))
			} else {
				dp[i] = max(dp[i-1], int64(nums[i]))
			}
		}
	}
	return dp[n-1]
}
