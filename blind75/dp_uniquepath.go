package blind

import "fmt"

// c(m+n-2,m-1)
func uniquePaths(m int, n int) int {
	N := m + n - 2
	M := m - 1
	return factorial(N) / (factorial(M) * factorial(N-M))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}

func dp_uniquePath(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 0
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	fmt.Println(dp)
	return dp[m-1][n-1]
}
