package blind

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minWindow(s string, t string) string {
	start, end := 0, 0
	tM := make(map[rune]int)
	var contained func() bool
	res := ""

	for _, tc := range t {
		tM[tc] += 1
	}

	contained = func() bool {
		for _, v := range tM {
			if v > 0 {
				return false
			}
		}
		return true
	}

	// "ADOBECODEBANC", "ABC"
	flag := false
	for start <= end && end < len(s) {
		if _, exists := tM[rune(s[end])]; exists && !flag {
			tM[rune(s[end])] -= 1
		}

		if contained() {
			flag = true
			println(start, end)
			if len(res) == 0 {
				res = s[start : end+1]
			} else {
				if end-start+1 < len(res) {
					res = s[start : end+1]
				}
			}

			if _, exists := tM[rune(s[start])]; exists {
				tM[rune(s[start])] += 1
			}
			start += 1
			continue
		}

		end += 1
	}

	return res
}
