package blind

import "fmt"

func maxF(a, b, c int) int {
	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	return m
}

func minF(a, b, c int) int {
	m := a

	if b < m {
		m = b
	}

	if c < m {
		m = c
	}

	return m
}

// 输入: nums = [2,3,-2,4]
func maxProduct(nums []int) int {
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dpMax[i] = maxF(nums[i]*dpMax[i-1], nums[i]*dpMin[i-1], nums[i])
		dpMin[i] = minF(nums[i]*dpMax[i-1], nums[i]*dpMin[i-1], nums[i])
	}

	fmt.Println(dpMax, dpMin)
	maxV := dpMax[0]
	for i := 0; i < len(nums); i++ {
		if dpMax[i] > maxV {
			maxV = dpMax[i]
		}
	}
	return maxV
}

// dp[i] = max{dp[i-1] + nums[i],nums[i]}
func maxSubArrayDp(nums []int) int {
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
