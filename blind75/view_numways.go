package blind

func numWays(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	return numWays(n-1) + numWays(n-2)
}

func numWaysIter(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func numWaysIterV(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	pre, curr := 1, 2
	for i := 3; i <= n; i++ {
		pre, curr = curr, pre+curr
	}

	return curr
}
