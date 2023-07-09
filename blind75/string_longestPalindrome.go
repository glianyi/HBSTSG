package blind

// cbbd
func longestPalindrome(s string) string {
	if len(s) == 1 || len(s) == 0 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	start, maxLen := -1, 0

	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	for l := 2; l < len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := i + l - 1
			if j >= len(s) {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if l == 2 {
					dp[i][j] = true
				}
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] && l > maxLen {
				maxLen = l
				start = i
			}
		}
	}

	return s[start : start+maxLen]
}

func countSubstrings(s string) int {
	count := 0
	if len(s) == 0 {
		return count
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		count += 1
	}

	for l := 2; l < len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := i + l - 1
			if j >= len(s) {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if l == 2 {
					dp[i][j] = true
				}
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] {
				count += 1
			}
		}
	}
	return count
}
