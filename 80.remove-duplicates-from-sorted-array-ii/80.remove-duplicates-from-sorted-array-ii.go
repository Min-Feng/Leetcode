/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
package main

// 失敗例子, 保留當時思考的樣貌
func removeDuplicates(nums []int) int {
	maxRepeat := 2
	move := 0
	repeat := 1

	// 先思考 move index 什麼情境才需要移動 或 不移動
	// 如果 swap 的時候, 選擇把一整個 slice append 到 前面的 slice
	// 那麼 slice 長度會一直變動, 該如何處理?
	for i := 1; i < len(nums); {
		switch {
		case nums[move] != nums[i] && repeat <= maxRepeat:
			repeat = 1
			move++
			i++

		case nums[move] != nums[i] && repeat > maxRepeat:
			repeat = 1
			nums = append(nums[:move], nums[i:]...)
			i = move + 1

		case nums[move] == nums[i] && repeat <= maxRepeat:
			repeat++
			move++
			i++

		case nums[move] == nums[i] && repeat > maxRepeat:
			repeat++
			i++
		}
	}
	return move + 1
}

// 第一個版本 歸納條件整理的不夠好
// 以為有四個條件
// 但其實只有三個
func removeDuplicatesV2(nums []int) int {
	maxRepeat := 2
	move := 0
	repeat := 1

	for i := 1; i < len(nums); {
		switch {
		case repeat > maxRepeat:
			repeat = 1
			nums = append(nums[:move], nums[i+1:]...)

		case nums[move] != nums[i] && repeat <= maxRepeat:
			repeat = 1
			move++
			i++

		case nums[move] == nums[i] && repeat <= maxRepeat:
			repeat++
			move++
			i++
		}
	}
	return move + 1
}

// 以為只有三個條件, 其實依然是四個條件...
// 但後來又發現只有三個條件
// 這就是 事前沒想清楚
// 到底有哪些事件需要移動相關指標 及 進行交換的動作
//
// 與下題概念類似 但允許 repeat 的次數, 大大增加複雜性
// 同樣是使用兩個 index 前進
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicatesV3(nums []int) int {
	maxRepeat := 2
	repeat := 1

	for i := 1; i < len(nums); {
		switch {
		case nums[i-1] != nums[i]:
			repeat = 1
			i++

		case nums[i-1] == nums[i] && repeat < maxRepeat:
			repeat++
			i++

		case nums[i-1] == nums[i] && repeat >= maxRepeat:
			repeat++
			nums = append(nums[:i], nums[i+1:]...) // remove i element
		}
	}
	return len(nums)
}

// 一次比較兩個數字
func removeDuplicatesV4(nums []int) int {
	for i := 2; i < len(nums); {
		if nums[i-1] == nums[i] && nums[i-2] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...) // remove i element
		} else {
			i++
		}
	}
	return len(nums)
}

// 若改成寫頭是「可以寫的位置」，即為寫頭以前的是寫入過的，當前位置是要準備寫入的位置。
// 拿 index -2 位罝的數字當比較，中間是什麼就不重要了。
// 因為有序，若讀頭數字 = 寫頭-2的數字，則中間的數字一定要相同
// https://www.notion.so/80-Remove-Duplicates-from-Sorted-Array-II-M-3b11b56f08574893955367fbab22c94a
func removeDuplicatesV5(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	write := 2
	for read := 2; read < len(nums); read++ {
		// 拿 index -2 位罝的數字當比較
		if nums[write-2] != nums[read] {
			nums[write] = nums[read]
			write++
		}
	}
	return write
}

// @lc code=end
