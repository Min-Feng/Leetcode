package my_sol_1

func longestCommonPrefix(strs []string) string {
	var returnPrefix string

	// 要考慮無任何輸入的情況
	if len(strs) == 0 {
		return returnPrefix
	}

loop:
	for i := range strs[0] {
		prefix := strs[0][:i+1]
		for j, str := range strs {
			// 要考慮取樣的prefix字串大於陣列中的其他字串
			if len(prefix) > len(str) {
				break loop
			}

			if str[:i+1] != prefix {
				break loop
			}

			// 執行到最後一個項目應該用index判斷
			// 不應該用字串內容 查看是否相等 來決定
			// 若陣列中其他陣列的字串內容 與 最後一個字串相同 會造成錯誤
			if j == len(strs)-1 {
				returnPrefix = prefix
			}
		}
	}
	return returnPrefix
}
