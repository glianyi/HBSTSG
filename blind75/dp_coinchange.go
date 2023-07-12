package blind

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c >= 0 && dp[i-c] != -1 {
				if dp[i] == -1 {
					dp[i] = dp[i-c] + 1
					continue
				}
				if dp[i] != -1 && dp[i-c]+1 < dp[i] {
					dp[i] = dp[i-c] + 1
				}
			}
		}
	}
	return dp[amount]
}
