// Given an array of integers nums which is sorted in ascending order, and an int
// eger target, write a function to search target in nums. If target exists, then r
// eturn its index. Otherwise, return -1.
//
//
// Example 1:
//
//
// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4
//
//
// Example 2:
//
//
// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 104
// -9999 <= nums[i], target <= 9999
// All the integers in nums are unique.
// nums is sorted in an ascending order.
//
// Related Topics Binary Search
// 👍 1326 👎 62
package main

import (
	"fmt"
)

func main() {
	r := searchRecursion([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println("result", r)
}

func searchRecursion(nums []int, target int) (result int) {
	if len(nums) == 0 {
		return -1
	}

	head := 0
	tail := len(nums) - 1
	pivot := head + (tail-head)/2

	switch {
	case nums[pivot] == target:
		return pivot

	case nums[pivot] > target:
		return searchRecursion(nums[:pivot], target)

	case nums[pivot] < target:
		result = searchRecursion(nums[pivot+1:], target)
		if result != -1 {
			compensate := pivot + 1
			return compensate + result
		}
		return result
	}

	return
}
