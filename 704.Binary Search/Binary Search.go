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

import "fmt"

func main() {
	r := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	fmt.Println("result", r)
}

// a.
// 取左右極端值, 並選擇一個基準點
// 這裡選擇中間點 (high + low) / 2, 但不知道要不要根據餘數+1
//
// b. 比較基準點 和 目標值
// b-1: 基準點 > 目標值, 讓 右極端值=基準點值
// b-2: 基準點 < 目標值, 讓 左極端值=基準點值
//
// c.
// 重新計算基準點 思考不出
// 計算基準點的正確邏輯是什麼, 只能 try
//
// d.
// 程式失敗, 無法通過所有情況
//
// e.
// 看完解答, 疑惑 low 有可能 大於 high? 想不通
// 我的迴圈中止條件, 稿的自己很亂

// 解題時的 困難點
// 1. 目標數字不存在時的中止條件
// 2. pivot 的定義是? 單純 pivot := (high + low) / 2 會進入無窮迴圈
//
// 3.
// right = pivot - 1
// left = pivot + 1
// +1 -1 的運作, 真的有辦法 直接靠思考得知嗎
// 還是只能用觀察法判斷, 可能進入迴圈 才事後補償
// 我開始只有單純 right = pivot, left = pivot

// 釐清思考方向
// right = pivot - 1
// left = pivot + 1
// 因為 前面有先判斷 目標值 和 基準點數值是否相等
// 不相等才會進入 移動 right left 的條件式
// 判斷過是否相等 所以不需要使用 right = pivot 這種 pivot 策略
// 直接加減1 進一步縮小範圍就好
// 由於有 加減1 縮小範圍的動作, 所以 中止條件 left 是會大於 right 的
//
// 若是依照原本的 right = pivot 把判斷過的條件又包含在下次迴圈
// 那等於沒進行判斷動作, 重複做一樣的事情!

// 注意大數 相加, 改用 pivot = low + ((high - low) / 2)
// https://ai.googleblog.com/2006/06/extra-extra-read-all-about-it-nearly.html

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	low := start
	high := end

	result := -1
	pivot := (high + low) / 2
	count := make(map[int]int)

Loop:
	for pivot >= start && pivot <= end && count[low] < 2 && count[high] < 2 {

		switch {
		case nums[pivot] == target:
			result = pivot
			break Loop

		case nums[pivot] < target:
			low = pivot
			count[low]++

			remainder := (high + low) % 2
			if remainder == 0 {
				pivot = (high + low) / 2
			} else {
				pivot = (high+low)/2 + 1
			}

		case nums[pivot] > target:
			high = pivot
			count[high]++

			remainder := (high + low) % 2
			if remainder == 0 {
				pivot = (high + low) / 2
			} else {
				pivot = (high+low)/2 - 1
			}

		}
	}
	return result
}

// leetcode submit region end(Prohibit modification and deletion)
