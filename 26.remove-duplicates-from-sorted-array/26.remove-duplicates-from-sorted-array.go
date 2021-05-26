/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1
	changeIndex := 0
	for _, num := range nums {
		repeatN := nums[changeIndex]
		if repeatN != num {
			count++
			changeIndex++
			nums[changeIndex] = num
		}
	}
	return count
}

// @lc code=end
