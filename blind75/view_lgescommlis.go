package blind

// text1 = "abcde", text2 = "ace" 3

// dp[i,j] = t1[0:i] t2[0,j]的最长公共子序列
// dp[i,j] = dp[i-1,j-1] + i if t1[i] == t2[j]
// dp[i,j] = max{dp[i-1][j],dp[j-1][i]} if t1[i] != t2[j]

func longestCommonSubsequenceView(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 0; i < len(text1)+1; i++ {
		dp[i][0] = 0
	}
	for j := 0; j < len(text2)+1; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}
