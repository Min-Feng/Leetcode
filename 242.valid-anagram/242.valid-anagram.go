/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start

package main

import "strings"

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
// Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Valid Anagram.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	record := [26]byte{}

	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
		record[t[i]-'a']--
	}

	var total int
	for _, n := range record {
		total += int(n)
	}

	return total == 0
}

// 雖然速度慢, 但想法挺有趣
//
// Runtime: 272 ms, faster than 7.34% of Go online submissions for Valid Anagram.
// Memory Usage: 8.7 MB, less than 5.03% of Go online submissions for Valid Anagram.
func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	for i := range s {
		subString := s[i : i+1]
		t = strings.Replace(t, subString, "", 1)
	}

	return t == ""
}

// @lc code=end
