/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

package main

// https://leetcode.com/problems/group-anagrams/
func main() {

}

func groupAnagrams(strs []string) [][]string {
	list := make([][]string, 0)

	// key is character count collection, value is index of list
	record := make(map[[26]int]int)

	for _, str := range strs {
		charCount := [26]int{}

		for i := range str {
			char := str[i] - 'a'
			charCount[char]++
		}

		index, ok := record[charCount]
		if !ok {
			list = append(list, []string{str})
			record[charCount] = len(list) - 1
		} else {
			list[index] = append(list[index], str)
		}
	}

	return list
}
