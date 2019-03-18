package my_sol_1

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	bracketPair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, r := range s {
		if isLeftBracket(r) {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop != bracketPair[r] {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}

func isLeftBracket(r rune) bool {
	if r == '(' || r == '[' || r == '{' {
		return true
	}
	return false
}
