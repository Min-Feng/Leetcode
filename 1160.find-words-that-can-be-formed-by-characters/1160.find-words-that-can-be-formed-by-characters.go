/*
 * @lc app=leetcode id=1160 lang=golang
 *
 * [1160] Find Words That Can Be Formed by Characters
 */

// @lc code=start
package main

func main() {

}

// Runtime: 72 ms, faster than 26.25% of Go online submissions for Find Words That Can Be Formed by Characters.
// Memory Usage: 6.6 MB, less than 7.50% of Go online submissions for Find Words That Can Be Formed by Characters.
func countCharactersV1(words []string, chars string) int {
	target := make(map[rune]int)
	for _, ch := range chars {
		target[ch]++
	}

	var sum int
	for i := 0; i < len(words); i++ {
		sample := make(map[rune]int)
		for _, ch := range words[i] {
			sample[ch]++
		}
		if len(target) < len(sample) {
			continue
		}

		for key, value := range sample {
			if target[key] >= value {
				delete(sample, key)
			}
		}
		if len(sample) == 0 {
			sum += len(words[i])
		}
	}

	return sum
}

// 將 v2 兩個 loop 改成 一個
func countCharactersV2(words []string, chars string) int {
	target := make([]byte, 26)
	for _, ch := range []byte(chars) {
		target[ch-'a']++
	}

	var sum int
	for i := 0; i < len(words); i++ {
		sample := make([]byte, 26)
		var count byte
		for _, ch := range []byte(words[i]) {
			index := ch - 'a'
			sample[index]++
			count++
			if target[index] < sample[index] {
				count = 0
				break
			}
		}
		sum += int(count)
	}

	return sum
}

func countCharactersV3(words []string, chars string) int {
	target := make([]byte, 26)
	for _, ch := range []byte(chars) {
		target[ch-'a']++
	}

	var sum int
	for i := 0; i < len(words); i++ {
		sample := make([]byte, 26)
		var count byte
		for _, ch := range []byte(words[i]) {
			index := ch - 'a'
			sample[index]++
			count++
			if target[index] < sample[index] {
				count = 0
				break
			}
		}
		sum += int(count)
	}

	return sum
}
