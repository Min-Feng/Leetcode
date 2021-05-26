/*
 * @lc app=leetcode id=961 lang=golang
 *
 * [961] N-Repeated Element in Size 2N Array
 */

// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/

// @lc code=start
package main

// Runtime: 20 ms, faster than 100.00% of Go online submissions for N-Repeated Element in Size 2N Array.
// Memory Usage: 6.8 MB, less than 31.58% of Go online submissions for N-Repeated Element in Size 2N Array.
func repeatedNTimes(nums []int) int {
	record := make(map[int]bool)

	for _, num := range nums {
		if record[num] {
			return num
		}
		record[num] = true
	}

	return nums[0]
}

// Runtime: 24 ms, faster than 98.25% of Go online submissions for N-Repeated Element in Size 2N Array.
// Memory Usage: 6.9 MB, less than 31.58% of Go online submissions for N-Repeated Element in Size 2N Array.
func repeatedNTimesV2(nums []int) int {
	record := [10000]int{}

	for _, num := range nums {
		if record[num] != 0 {
			return num
		}
		record[num]++
	}

	return nums[0]
}

// @lc code=end
