package my_sol_1

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var stack []int

	for x != 0 {
		remaindeer := x % 10
		stack = append(stack, remaindeer)
		x = x / 10
	}

	for i, j := 0, len(stack)-1; j > i; i, j = i+1, j-1 {
		if stack[i] != stack[j] {
			return false
		}
	}
	return true
}
