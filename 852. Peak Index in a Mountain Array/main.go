// Let's call an array arr a mountain if the following properties hold:
//
//
// arr.length >= 3
// There exists some i with 0 < i < arr.length - 1 such that:
//
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
//
//
//
//
// Given an integer array arr that is guaranteed to be a mountain, return any i
// such that arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr
// .length - 1].
//
//
// Example 1:
// Input: arr = [0,1,0]
// Output: 1
// Example 2:
// Input: arr = [0,2,1,0]
// Output: 1
// Example 3:
// Input: arr = [0,10,5,2]
// Output: 1
// Example 4:
// Input: arr = [3,4,5,1]
// Output: 2
// Example 5:
// Input: arr = [24,69,100,99,79,78,67,36,26,19]
// Output: 2
//
//
// Constraints:
//
//
// 3 <= arr.length <= 104
// 0 <= arr[i] <= 106
// arr is guaranteed to be a mountain array.
//
//
//
// Follow up: Finding the O(n) is straightforward, could you find an O(log(n)) so
// lution? Related Topics Binary Search
// 👍 1079 👎 1368

package main

func peakIndexInMountainArrayLinearScan(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		next := i + 1
		if arr[i] > arr[next] {
			return i
		}
	}
	return len(arr) - 1
}

func peakIndexInMountainArrayBinarySearch(arr []int) int {
	start := 0
	head := start + 1

	last := len(arr) - 1
	tail := last - 1
	var pivot int

	for head <= tail {
		pivot = head + (tail-head)/2
		switch {
		case isPeak(arr, pivot):
			return pivot

		case isPivotUphill(arr, pivot):
			head = pivot + 1

		case isPivotDownhill(arr, pivot):
			tail = pivot - 1
		}
	}
	return pivot
}

func isPivotDownhill(arr []int, pivot int) bool {
	return arr[pivot] < arr[pivot-1] && arr[pivot] > arr[pivot+1]
}

func isPivotUphill(arr []int, pivot int) bool {
	return arr[pivot] > arr[pivot-1] && arr[pivot] < arr[pivot+1]
}

func isPeak(arr []int, pivot int) bool {
	return arr[pivot] > arr[pivot-1] && arr[pivot] > arr[pivot+1]
}
