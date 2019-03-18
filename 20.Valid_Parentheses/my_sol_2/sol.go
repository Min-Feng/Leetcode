package my_sol_2

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	bracketPair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, r := range s {
		if _, exist := bracketPair[r]; !exist {
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

	return len(stack) == 0
}
