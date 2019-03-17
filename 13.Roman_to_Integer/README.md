# 13. Roman to Integer

## 原題連結

<https://leetcode.com/problems/roman-to-integer/>

## My Solution

my_Sol_1：[Link](my_sol_1/my_sol_1.go)

**加強字串的使用方式**

1. 先用map型別建立對應的表
2. 對輸入字串進行for迴圈
3. 根據取得的字元得到對應數字,進行累加
4. 佔用的記憶體太多  
   在函數中make建立map型別每次呼叫都會重新建立,造成*時間浪費*,以及多付出記憶體  
   Memory Usage: 5 MB, less than 8.00% of Go online submissions for Roman to Integer.  

   改為在函數外部make,可改善情況  
   Memory Usage: 3 MB, less than 40.00% of Go online submissions for Roman to Integer.

## Website Provide Sol：

Sol_1：[Link](website_sol_1/sol_1.go)

1. 方法類似我的想到解法,但是將型別map從`map[string]int`改為 `map[byte]int`
2. 對一個字串使用range,s[i]是byte型別,v是rune型別
3. 遇到特殊情況,使用減法,而不是從map建立好預設字串該有的動作


## Reference
