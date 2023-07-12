package blind

import "fmt"

// [1,2,3,1]
// 输出：4

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i < len(nums)+1; i++ {
		dp[i] = maxInt(dp[i-2]+nums[i-1], dp[i-1])
	}
	fmt.Println(dp)
	return dp[len(nums)]
}
