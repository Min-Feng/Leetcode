/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

package main

import "sort"

// https://leetcode.com/problems/top-k-frequent-elements/

// Runtime: 12 ms, faster than 90.72% of Go online submissions for Top K Frequent Elements.
// Memory Usage: 5.4 MB, less than 94.16% of Go online submissions for Top K Frequent Elements.
//
// map結構, 對出現次數進行統計
// 然後使用套件 sort 求得最高頻率
func topKFrequent(nums []int, k int) []int {
	stats := make(map[int]int)
	for _, num := range nums {
		stats[num]++
	}

	type record struct {
		num   int
		count int
	}
	list := make([]record, 0, len(stats))
	for num, count := range stats {
		list = append(list, record{
			num:   num,
			count: count,
		})
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].count > list[j].count
	})

	result := make([]int, k)
	for i := range result {
		result[i] = list[i].num
	}

	return result
}
