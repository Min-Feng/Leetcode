// Given an integer n, return a string array answer (1-indexed) where:
//
//
// answer[i] == "fizzbuzz" if i is divisible by 3 and 5.
// answer[i] == "Fizz" if i is divisible by 3.
// answer[i] == "Buzz" if i is divisible by 5.
// answer[i] == i if non of the above conditions are true.
//
//
//
// Example 1:
// Input: n = 3
// Output: ["1","2","Fizz"]
// Example 2:
// Input: n = 5
// Output: ["1","2","Fizz","4","Buzz"]
// Example 3:
// Input: n = 15
// Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","1
// 3","14","FizzBuzz"]
//
//
// Constraints:
//
//
// 1 <= n <= 104
//
// 👍 1339 👎 1551

package leetcode

import (
	"strconv"
)

// leetcode submit region begin(Prohibit modification and deletion)

func fizzBuzz(n int) []string {
	result := make([]string, 0, n)
	var cond1, cond2 bool

	for i := 1; i <= n; i++ {
		cond1 = (i % 3) == 0
		cond2 = (i % 5) == 0

		switch {
		case cond1 && cond2:
			result = append(result, "FizzBuzz")

		case cond1:
			result = append(result, "Fizz")

		case cond2:
			result = append(result, "Buzz")

		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

// leetcode submit region end(Prohibit modification and deletion)
