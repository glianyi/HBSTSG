package blind

// AABABAA
func characterReplacement(s string, k int) int {
	start, maxLen, maxFreq := 0, 0, 0
	freq := make(map[rune]int)
	for i, c := range s {
		freq[c] += 1
		maxFreq = maxInt(maxFreq, freq[c])

		if i-start+1 > maxFreq+k {
			maxLen = maxFreq + k
			freq[rune(s[start])] -= 1
			start += 1
			continue
		}
		maxLen = maxInt(maxLen, i-start+1)
	}

	return maxLen
}
