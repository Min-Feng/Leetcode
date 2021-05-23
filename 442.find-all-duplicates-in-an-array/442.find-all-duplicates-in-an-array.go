/*
 * @lc app=leetcode id=442 lang=golang
 *
 * [442] Find All Duplicates in an Array
 */
package main

import "sort"

func main() {
	findDuplicatesCycleSort([]int{4, 3, 2, 7, 8, 2, 3, 1})
}

// Runtime: 68 ms, faster than 22.99% of Go online submissions for Find All Duplicates in an Array.
// Memory Usage: 8.8 MB, less than 6.90% of Go online submissions for Find All Duplicates in an Array.
func findDuplicates(nums []int) []int {
	result := make([]int, 0, len(nums))
	record := make(map[int]int)
	for _, num := range nums {
		record[num]++
		if record[num] == 2 {
			result = append(result, num)
		}
	}
	return result
}

// Runtime: 48 ms, faster than 95.40% of Go online submissions for Find All Duplicates in an Array.
// Memory Usage: 7.6 MB, less than 19.54% of Go online submissions for Find All Duplicates in an Array.
func findDuplicatesV2(nums []int) []int {
	result := make([]int, 0)
	record := make([]int, 100001)

	for _, num := range nums {
		record[num]++
		if record[num] == 2 {
			result = append(result, num)
		}
	}
	return result
}

// 看到限制 1次 2次 就要想到 使用 bitwise 進行紀錄
// 但這個是錯誤作法, uint64 只有 64個槽 可以紀錄, num 的數字遠超過
// 失敗的作法
func findDuplicatesV3(nums []int) []int {
	result := make([]int, 0)
	var record, index uint64

	for _, num := range nums {
		index = 1 << num
		record += index
		if record&index != index {
			record = record - index
			result = append(result, num)
		}
	}
	return result
}

// Runtime: 84 ms, faster than 5.75% of Go online submissions for Find All Duplicates in an Array.
// Memory Usage: 7 MB, less than 35.63% of Go online submissions for Find All Duplicates in an Array.
func findDuplicatesV4(nums []int) []int {
	result := make([]int, 0)
	sort.Ints(nums)

	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			result = append(result, nums[i])
			i += 2
		} else {
			i++
		}
	}
	return result
}

// https://haogroot.com/2020/11/26/cyclic_sort-leetcode/
// 主要想法: 每個數字應該恰好在自己的位置上
//
// Runtime: 48 ms, faster than 96.46% of Go online submissions for Find All Duplicates in an Array.
// Memory Usage: 6.8 MB, less than 89.38% of Go online submissions for Find All Duplicates in an Array.
func findDuplicatesCycleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; {
		position := nums[i] - 1 // 注意範圍是 0~n 或 1~n
		if nums[position] != nums[i] {
			nums[i], nums[position] = nums[position], nums[i]
		} else {
			i++
		}
	}

	var r []int
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			r = append(r, nums[i])
			// r = append(r, i) // 注意: 是其元素 沒有按照位置, index 都是唯一值
		}
	}
	return r
}
