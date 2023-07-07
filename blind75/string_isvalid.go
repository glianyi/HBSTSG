package blind

func isValid(s string) bool {
	stack := []rune{}

	match := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		}
		if c == ')' || c == ']' || c == '}' {
			if match[c] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
