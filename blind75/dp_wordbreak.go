package blind

// "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if _, exists := m[s[j:i]]; exists {
				dp[i] = dp[j]
				break
			}
		}
	}
	return dp[len(s)]
}
