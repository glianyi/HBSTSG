package blind

// nums = [10,9,2,5,3,7,101,18],4
// dp[0]=1 dp[1]=1 dp[2]=1 dp[3]=2,dp[4]=2,dp[5]=3,dp[6]=4,dp[7]=4
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	res := 0
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}
