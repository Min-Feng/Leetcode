package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(backtraceV8([]int{0, 1, 2}))
}

// 目標列出 n 位數的所有可能性
// 也就是每個位數 都要出現 0~9
//
// 第一個位數 會有 10 個情況,
// 第二個位數 會根據前一個位數 又產生 10 個情況

// 因為做不出來, 改成簡單的想法
// 先考慮 一個位數的情況 0~9 如何撰寫
func backtraceV1() {
	// 錯誤作法
	// result := make([]int, 9)
	// for i := 0; i < 9; i++ {
	// 	result[i] = i
	// }
	// fmt.Println(result)

	result := make([][]int, 10) // 0~9 是 10個數字
	for i := 0; i <= 9; i++ {
		num := make([]int, 1)
		num[0] = i // 之後多個位數, 從這裡開始延伸
		result[i] = num
	}

	fmt.Println(result)
}

// 再來思考兩個位數
// 實際執行後
// 除了記憶數值的問題 次數也不對 應該是 10*10 不是 10*2
func backtraceV2() {
	n := 2
	result := make([][]int, 0)
	for i := 0; i <= 9; i++ {
		var num []int
		for j := 0; j < n; j++ { // 事後討論: 位數填寫, 不能用迴圈, 要指定對應位數, 要給什麼數值, 迴圈做不到這個行為, 參考 V3 作法m
			num = make([]int, n)

			// 類似 對角線的感覺, 先把所有出現 1 的位置
			// 填寫上去, 但是這樣做沒有記憶到之前位數
			// 想不到如何記憶
			num[j] = i
			result = append(result, num)
		}
	}

	fmt.Println(result)
	fmt.Println(len(result))
}

// 參考網路範例 求得
// https://youtu.be/nrHTtjkYEyQ?t=136
func backtraceV3() {
	r := make([][]int, 0)

	// 記憶數值方法2: 一開始就 陣列宣告
	// 失敗作法
	// num := make([]int, 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if i != j && j != k && k != i {
					// 記憶數值
					// 藉由 迴圈指針 i j k 紀錄位置
					// 在最後一個迴圈 再把數值 填入到結果

					// 記憶數值方法1: 最後才進行 陣列宣告
					num := make([]int, 0, 3)
					num = append(num, i, j, k)
					// [[0 1 2] [0 2 1] [1 0 2] [1 2 0] [2 0 1] [2 1 0]]

					// 記憶數值方法2: 一開始就 陣列宣告
					// 會因為引用同樣的變數位置, 造成所有排列結果都相同
					// [[2 1 0] [2 1 0] [2 1 0] [2 1 0] [2 1 0] [2 1 0]]
					//
					// num[0] = i
					// num[1] = j
					// num[2] = k
					r = append(r, num)
				}
			}
		}
	}
	fmt.Println(r)
}

// 改善 V3 記憶數值方法2的錯誤
func backtraceV4() {
	r := make([][]int, 0)
	num := make([]int, 3)

	for i := 0; i < 3; i++ {
		num[0] = i
		for j := 0; j < 3; j++ {
			num[1] = j
			for k := 0; k < 3; k++ {
				num[2] = k
				if i != j && j != k && k != i {
					// 記憶數值
					// 藉由 迴圈指針 i j k 紀錄位置
					// 在最後一個迴圈 再把數值 填入到結果
					r = append(r, append(make([]int, 0, 3), num...))
				}
			}
		}
	}
	fmt.Println(r)
}

// 遞迴法 但沒有排除重複數字
func backtraceV5(nums []int) [][]int {
	target := make([]int, len(nums))
	return V5helper(nums, target, 0)
}

func V5helper(nums []int, target []int, pointer int) [][]int {
	if pointer == len(nums) {
		buf := make([]int, len(nums))
		copy(buf, target) //  重要
		return append(make([][]int, 0), buf)
	}

	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		target[pointer] = nums[i]
		r := V5helper(nums, target, pointer+1)
		result = append(result, r...)
	}

	return result
}

// 改進 V5 排除重複
func backtraceV6(nums []int) [][]int {
	target := make([]int, len(nums))
	return V6helper(nums, target, 0)
}

func V6helper(nums []int, target []int, pointer int) [][]int {
	if pointer == len(nums) {
		check := make([]int, len(nums))
		copy(check, target) // 如果沒用 copy, sort 會影響 呼叫端的數值
		sort.Ints(check)
		for i := 0; i < len(check)-1; i++ {
			if check[i] == check[i+1] {
				return make([][]int, 0)
			}
		}

		buf := make([]int, len(nums))
		copy(buf, target) // 如果沒用 copy , 最後所有結果, 都是相同的
		return append(make([][]int, 0, 1), buf)
	}

	result := make([][]int, 0)
	// for i := pointer; i < len(nums); i++ { // 錯誤 過濾: [[0 1 2] [0 2 2] [1 1 2] [1 2 2] [2 1 2] [2 2 2]]
	for i := 0; i < len(nums); i++ {
		target[pointer] = nums[i]
		r := V6helper(nums, target, pointer+1)
		result = append(result, r...)
	}

	return result
}

func newBacktraceV8Helper(nums []int) *backtraceV8Helper {
	return &backtraceV8Helper{
		maxQuantity: len(nums),
		source:      nums,
		target:      make([]int, len(nums)),
		use:         make([]bool, len(nums)),
	}
}

type backtraceV8Helper struct {
	maxQuantity int
	source      []int
	target      []int
	use         []bool

	results [][]int
}

func (v8 *backtraceV8Helper) run(currentIndex int) {
	if currentIndex == v8.maxQuantity {
		result := make([]int, v8.maxQuantity)
		copy(result, v8.target)
		v8.results = append(v8.results, result)
		return
	}

	// 迴圈的 i 固定都從 0 開始, 不因 currentIndex 有所改變
	// 改成另外用一個記憶體位置 use 來紀錄哪些條件已經使用過, 直接跳過
	for i := 0; i < v8.maxQuantity; i++ {
		if v8.use[i] {
			continue
		}

		v8.target[currentIndex] = v8.source[i]
		v8.use[i] = true // 訣竅 標記是否用過, 以便 run迴圈執行時, 可以跳過不重複執行
		v8.run(currentIndex + 1)
		v8.use[i] = false // 重點: 在run遞迴已經用了 ues, 為了下個迴圈使用, 要標示為取消
	}
	return
}

func (v8 *backtraceV8Helper) Results() [][]int {
	return v8.results
}

func backtraceV8(nums []int) [][]int {
	helper := newBacktraceV8Helper(nums)
	helper.run(0)
	return helper.Results()
}

// 嘗試更簡潔的寫法, 但編譯失敗: use of [...] array outside of array literal
// func backtraceV7(nums []int) [][...]int {
// 	var target [...]int
// 	copy(target[:], nums)
// 	return V7helper(nums, target, 0)
// }
//
// func V7helper(nums []int, target [...]int, pointer int) [][...]int {
// 	if pointer == len(nums) {
// 		sort.Ints(target[:])
// 		for i := 0; i < len(target)-1; i++ {
// 			if target[i] == target[i+1] {
// 				return make([][...]int, 0)
// 			}
// 		}
// 		return append(make([][...]int, 0, 1), target)
// 	}
//
// 	result := make([][...]int, 0)
// 	// for i := pointer; i < len(nums); i++ { // 錯誤 過濾: [[0 1 2] [0 2 2] [1 1 2] [1 2 2] [2 1 2] [2 2 2]]
// 	for i := 0; i < len(nums); i++ {
// 		target[pointer] = nums[i]
// 		r := V7helper(nums, target, pointer+1)
// 		result = append(result, r...)
// 	}
//
// 	return result
// }
