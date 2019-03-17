package my_sol_1

var romanKey = make(map[string]int)

func romanToInt(s string) int {
	//改為在函數外部make,避免每次呼叫都重新建立map
	//var romanKey = make(map[string]int)
	makeRomanKeyData()

	var nowKey string
	var nextKey string
	var num int
	for i := 0; len(s) > i; i++ {
		nowKey = s[i : i+1]
		if i != len(s)-1 {
			nextKey = s[i+1 : i+2]
			if romanKey[nowKey] < romanKey[nextKey] {
				nowKey = s[i : i+2]
				i++ //若本次取樣字串為兩個字元一組,則跳過下次的循環計算
			}
		}
		num += romanKey[nowKey]
	}
	return num
}

func makeRomanKeyData() {
	romanKey["I"] = 1
	romanKey["IV"] = 4
	romanKey["IX"] = 9
	romanKey["V"] = 5
	romanKey["X"] = 10
	romanKey["XL"] = 40
	romanKey["XC"] = 90
	romanKey["L"] = 50
	romanKey["C"] = 100
	romanKey["CD"] = 400
	romanKey["CM"] = 900
	romanKey["D"] = 500
	romanKey["M"] = 1000
}
