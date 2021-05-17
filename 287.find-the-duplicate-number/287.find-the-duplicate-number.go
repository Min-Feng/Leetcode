/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
package main

import "sort"

func findDuplicate(nums []int) int {
	mark := make(map[int]int)
	for _, num := range nums {
		mark[num]++
		if mark[num] != 1 {
			return num
		}
	}
	return 0
}

// Time Limit Exceeded
func findDuplicateV1(nums []int) int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return 0
}

func findDuplicateV2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i+1 < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return 0
}

// https://medium.com/@orionssl/%E6%8E%A2%E7%B4%A2-floyd-cycle-detection-algorithm-934cdd05beb9
func findDuplicateV3(nums []int) int {
	return 0
}

// @lc code=end
