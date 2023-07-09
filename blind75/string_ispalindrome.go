package blind

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for left, right := 0, len(s)-1; left <= right; {
		for !unicode.IsLetter(rune(s[left])) {
			left += 1
		}
		for !unicode.IsLetter(rune(s[right])) {
			right -= 1
		}

		if s[left] == s[right] {
			left += 1
			right -= 1
		} else {
			return false
		}
	}

	return true
}
