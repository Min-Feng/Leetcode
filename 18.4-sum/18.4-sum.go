/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/4sum/

// https://www.notion.so/18-4Sum-d4fb5b780ae24535ae4d545ed4f5261e

func main() {
	// result := fourSumV2([]int{-500, -481, -480, -469, -437, -423, -408, -403, -397, -381, -379, -377, -353, -347, -337, -327, -313, -307, -299, -278, -265, -258, -235, -227, -225, -193, -192, -177, -176, -173, -170, -164, -162, -157, -147, -118, -115, -83, -64, -46, -36, -35, -11, 0, 0, 33, 40, 51, 54, 74, 93, 101, 104, 105, 112, 112, 116, 129, 133, 146, 152, 157, 158, 166, 177, 183, 186, 220, 263, 273, 320, 328, 332, 356, 357, 363, 372, 397, 399, 420, 422, 429, 433, 451, 464, 484, 485, 498, 499}, 2139)
	result := fourSumV5([]int{-497, -494, -484, -477, -453, -453, -444, -442, -428, -420, -401, -393, -392, -381, -357, -357, -327, -323, -306, -285, -284, -263, -262, -254, -243, -234, -208, -170, -166, -162, -158, -136, -133, -130, -119, -114, -101, -100, -86, -66, -65, -6, 1, 3, 4, 11, 69, 77, 78, 107, 108, 108, 121, 123, 136, 137, 151, 153, 155, 166, 170, 175, 179, 211, 230, 251, 255, 266, 288, 306, 308, 310, 314, 321, 322, 331, 333, 334, 347, 349, 356, 357, 360, 361, 361, 367, 375, 378, 387, 387, 408, 414, 421, 435, 439, 440, 441, 470, 492}, 1682)
	// result := fourSumV2([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(result)
}

func fourSum(nums []int, target int) [][]int {
	h := helper{
		maxCount: 4,
		total:    len(nums),
		nums:     nums,
		target:   target,
		result:   make([][]int, 0),
	}
	use := make([]bool, len(nums))
	h.run(use, 0)
	return h.result
}

func fourSumV2(nums []int, target int) [][]int {
	h := helper{
		maxCount: 4,
		total:    len(nums),
		nums:     nums,
		target:   target,
		result:   make([][]int, 0),
	}
	use := make(map[int]int)
	h.runV2(use, 0)
	return h.result
}

type helper struct {
	maxCount int
	total    int
	nums     []int
	target   int
	result   [][]int
}

func (h *helper) run(use []bool, index int) {
	if h.count(use) > h.maxCount {
		return
	}

	if index == h.total {
		if h.count(use) != h.maxCount {
			return
		}

		sum := 0
		r := make([]int, 0, h.maxCount)

		for i, value := range use {
			if value == true {
				sum += h.nums[i]
				r = append(r, h.nums[i])
			}
		}

		if sum == h.target {
			sort.Ints(r)
			if h.matchResult(r) {
				return
			}
			h.result = append(h.result, r)
		}
		return
	}

	use[index] = true
	h.run(use, index+1)

	use[index] = false
	h.run(use, index+1)
}

func (h *helper) runV2(use map[int]int, index int) {
	if len(use) > h.maxCount {
		return
	}

	if index == h.total {
		if len(use) != h.maxCount {
			return
		}

		sum := 0
		r := make([]int, 0, h.maxCount)

		for _, value := range use {
			sum += value
			r = append(r, value) // v2 use variable is map, so range value is random
		}

		if sum == h.target {
			sort.Ints(r)
			if h.matchResult(r) {
				return
			}
			h.result = append(h.result, r)
		}
		return
	}

	use[index] = h.nums[index]
	h.runV2(use, index+1)

	delete(use, index)
	h.runV2(use, index+1)
}

func (h *helper) matchResult(r []int) bool {
	for _, list := range h.result {
		same := 0
		for i, v := range list {
			if r[i] == v {
				same++
			}
		}
		if same == h.maxCount {
			return true
		}
	}
	return false
}

func (h helper) count(use []bool) int {
	count := 0

	for _, value := range use {
		if value == true {
			count++
		}
	}

	return count
}

// Runtime: 136 ms, faster than 11.11% of Go online submissions for 4Sum.
// Memory Usage: 3 MB, less than 40.00% of Go online submissions for 4Sum.
func fourSumV3(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	n := len(nums)

	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for k := j + 1; k < n-1; k++ {
				for g := k + 1; g < n; g++ {

					sum := nums[i] + nums[j] + nums[k] + nums[g]
					if sum == target {
						r := []int{nums[i], nums[j], nums[k], nums[g]}
						if !isDuplicate(result, r) {
							result = append(result, r)
						}
					}

				}
			}
		}
	}
	return result
}

func isDuplicate(result [][]int, subResult []int) bool {
	for _, list := range result {
		same := 0
		for i, v := range list {
			if subResult[i] == v {
				same++
			}
		}
		if same == 4 {
			return true
		}
	}
	return false
}

// Runtime: 136 ms, faster than 11.11% of Go online submissions for 4Sum.
// Memory Usage: 3.1 MB, less than 32.89% of Go online submissions for 4Sum.
func fourSumV4(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	record := make(map[[4]int]bool)
	n := len(nums)

	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for k := j + 1; k < n-1; k++ {
				for g := k + 1; g < n; g++ {

					sum := nums[i] + nums[j] + nums[k] + nums[g]
					if sum == target {
						r := [4]int{nums[i], nums[j], nums[k], nums[g]}
						if !record[r] {
							result = append(result, r[:])
							record[r] = true
						}
					}

				}
			}
		}
	}
	return result
}

// Runtime: 196 ms, faster than 8.89% of Go online submissions for 4Sum.
// Memory Usage: 4 MB, less than 21.33% of Go online submissions for 4Sum.
//
// 只能對付 4sum
func fourSumV5(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	n := len(nums)
	record := make(map[[4]int]bool)

	// two sum variable
	var subTarget, start int
	var num, remainder int
	subRecord := make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {

			// two sum
			subTarget = target - nums[i] - nums[j]
			start = j + 1
			for k := start; k < n; k++ {
				num = nums[k]
				remainder = subTarget - num
				if subRecord[remainder] {
					r := [4]int{nums[i], nums[j], remainder, num}
					if !record[r] {
						result = append(result, r[:])
						record[r] = true
					}
				} else {
					subRecord[num] = true
				}
			}

			for key := range subRecord {
				delete(subRecord, key)
			}

		}
	}
	return result
}

// 在 V5 版本, 若把 two sum 抽成函數放在外面呼叫
// 會因為 反覆建立 map 結構, 導致 效能大幅降低
func twoSum(nums []int, target int, start int) [][]int {
	n := len(nums)
	stat := make(map[int]bool)
	var current, remainder int
	var result [][]int

	for k := start; k < n; k++ {
		current = nums[k]
		remainder = target - current
		if stat[remainder] {
			// 可能存在多個解, 不能直接回傳
			// return []int{remainder, current}
			result = append(result, []int{remainder, current})
		} else {
			stat[current] = true
		}
	}

	return result
}
