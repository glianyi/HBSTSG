package blind

// nums = [-2,1,-3,4,-1,2,1,-5,4] 连续子数组 [4,-1,2,1] 的和最大为6
// dp[i] = max(dp[i],dp[i-1]+nums[i])
// max(dp[])
func maxSum(nums []int) int {
	acc := nums[0]
	maxV := acc
	for i := 1; i < len(nums); i++ {
		if nums[i] > acc+nums[i] {
			acc = nums[i]
		} else {
			acc += nums[i]
		}

		if acc > maxV {
			maxV = acc
		}
	}
	return maxV
}

func maxSumDp(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-1]+nums[i], nums[i])
	}

	maxV := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > maxV {
			maxV = dp[i]
		}
	}

	return maxV
}
