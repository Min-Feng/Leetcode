/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

// @lc code=start

// https://leetcode.com/problems/single-number-iii/
package main

import "sort"

// 感覺類似 80. Remove Duplicates from Sorted Array II
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
//
// 第一次 submit 失敗
// 放棄這個版本
func singleNumber(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}

	sort.Ints(nums)

	// 因為題目說 最少兩個數字出現一次
	// 其他數字出現兩次
	// slice 一定為 2 的倍數
	// 以 2,3,4 個 為一個單位 進行判斷 ???

	var r []int
Loop:
	for i := 0; i < len(nums); {
		switch {
		case i+1 == len(nums): // 情境 0 1 1 2, 表示 i 在最後一個位置
			r = append(r, nums[i])
			i++

		case nums[i] == nums[i+1]: // 情境 0 0 2 2
			i += 2

			//  input=[1,2,1,3,2,5], runtime error: index out of range [6] with length 6
		case nums[i] != nums[i+1] && nums[i+1] == nums[i+2]: // 情境 0 1 1 2
			r = append(r, nums[i])
			if len(r) == 2 {
				break Loop
			}
			i += 3

		case nums[i] != nums[i+1] && nums[i+1] != nums[i+2]: // 情境 0 1 2 2
			r = append(r, nums[i], nums[i+1])
			break Loop
		}
	}
	return r
}

func singleNumberV2(nums []int) []int {
	r = r[:0]
	sort.Ints(nums)
	helperV2(nums)
	return r
}

var r []int

func helperV2(nums []int) {
	if len(r) >= 2 {
		return
	}
	if len(nums) == 0 {
		return
	}
	if len(nums) == 1 {
		r = append(r, nums[0])
		return
	}
	if len(nums) == 2 {
		r = append(r, nums[0], nums[1])
		return
	}

	switch {
	case nums[0] == nums[1]: // 情境 0 0 2 2
		helperV2(nums[2:])

	case nums[0] != nums[1] && nums[1] == nums[2]: // 情境 0 1 1 2
		r = append(r, nums[0])
		helperV2(nums[3:])

	case nums[0] != nums[1] && nums[1] != nums[2]: // 情境 0 1 2 2
		r = append(r, nums[0], nums[1])
		helperV2(nums[2:])
	}
}

// 改為 loop 版本
func singleNumberV3(nums []int) []int {
	var r []int
	sort.Ints(nums)
	length := len(nums)

	for i := 0; i < length; {
		if len(r) >= 2 {
			break
		}

		switch {
		case (length - i) == 1:
			r = append(r, nums[i])
			i++

		case (length - i) == 2:
			r = append(r, nums[i], nums[i+1])
			i += 2

		default:
			switch {
			case nums[i] == nums[i+1]: // 情境 0 0 2 2
				i += 2

			case nums[i] != nums[i+1] && nums[i+1] == nums[i+2]: // 情境 0 1 1 2
				r = append(r, nums[i])
				i += 3

			case nums[i] != nums[i+1] && nums[i+1] != nums[i+2]: // 情境 0 1 2 2
				r = append(r, nums[i], nums[i+1])
				i += 2
			}
		}
	}
	return r
}

// xor 版本
// https://bear-1111.medium.com/260-single-number-iii-d35b7cc0e0aa
func singleNumberV4(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	// 負號 2補數 = 1補數+1
	// bit clear 0011 &^ 0101
	// xor bit = 1 代表兩個數字不同的地方 或者說 標示
	// 想辦法找出哪個 bit 不同, 這裡選擇找最低 bit 為 1 當作 flag
	diffFlag := xor &^ (xor - 1)

	r := make([]int, 2)
	for _, n := range nums {
		if n&diffFlag == 0 {
			r[0] ^= n
		} else {
			r[1] ^= n
		}
	}
	return r
}

// @lc code=end
