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

func findDuplicateFloydCycle(nums []int) int {
	var EncounterIndex int
	var slow, fast int // index 移動速度差異

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			slow = nums[0] // 其值是指向下一個節點的 index

			fast = nums[0]
			fast = nums[fast]
		} else {
			slow = nums[slow]

			fast = nums[fast]
			fast = nums[fast]
		}

		if fast == slow {
			EncounterIndex = slow // 仔細想想
			break
		}
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			fast = nums[0]
			slow = nums[EncounterIndex]
		} else {
			slow = nums[slow]
			fast = nums[fast]
		}

		if fast == slow {
			return fast // 仔細想想
		}
	}
	return -1
}

func findDuplicateFloydCycleV2(nums []int) int {
	// 看到很多 node 指向同一個 node
	// 會猜 是否 link list 存在循環
	//
	// 又題目 範圍 限制範圍, 所以肯定 存在循環
	// nums containing n + 1 integers where each integer is in the range [1, n]
	start := 0

	// 設置起點位置
	fast := start
	slow := start

	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow {
			break
		}
	}

	// reset position
	fast = start

	for {
		fast = nums[fast]
		slow = nums[slow]
		if fast == slow {
			return fast
		}
	}
}

// @lc code=end
