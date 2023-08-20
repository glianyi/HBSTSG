package blind

// [3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列

// dp[i] = max(dp[0 <= j < i]) + 1 必须包含i元素
func lengthOfLISView(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		maxL := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > maxL {
				maxL = dp[j] + 1
			}
		}
		dp[i] = maxL
	}

	maxL := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > maxL {
			maxL = dp[i]
		}
	}

	return maxL
}
