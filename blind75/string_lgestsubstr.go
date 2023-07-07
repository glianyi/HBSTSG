package blind

// abcabcbb
func lengthOfLongestSubstring(s string) int {
	mark := make(map[rune]int)
	maxLen := 0
	start := 0
	for i, c := range s {
		if index, exists := mark[c]; exists {
			if maxLen < i-start+1 {
				maxLen = i - start
			}
			start = index + 1
			mark[c] = i
			continue
		}
		if maxLen < i-start+1 {
			maxLen = i - start + 1
		}
		mark[c] = i
	}
	return maxLen
}
