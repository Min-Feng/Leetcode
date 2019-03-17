package sol_1

// 對一個字串使用range,s[i]是byte型別,v是rune型別
// 用range無法給予下次的v是多少,所以只能用s[i]的形式,故使用byte型別
var romanKey = make(map[byte]int)

func romanToInt(s string) int {
	insertRomanKeyData()

	var num int
	for i := range s {
		if i != len(s)-1 && romanKey[s[i]] < romanKey[s[i+1]] {
			num -= romanKey[s[i]]
			continue
		}
		num += romanKey[s[i]]
	}
	return num
}

func insertRomanKeyData() {
	romanKey['I'] = 1
	romanKey['V'] = 5
	romanKey['X'] = 10
	romanKey['L'] = 50
	romanKey['C'] = 100
	romanKey['D'] = 500
	romanKey['M'] = 1000
}
