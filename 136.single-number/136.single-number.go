/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
package main

func singleNumber(nums []int) int {
	xor := 0
	for i := range nums {
		xor = xor ^ nums[i]
	}
	return xor
}

// @lc code=end
