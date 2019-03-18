# 14. Longest Common Prefix

## 原題連結

<https://leetcode.com/problems/longest-common-prefix/>

## My Solution

Sol_1：[Link](my_sol_1/sol_1.go)
實做中有幾點沒考慮到

1. 無任何輸入的情況
2. 取樣的prefix字串大於陣列中的其他字串
3. 執行到最後一個項目應該用index判斷

## Website Provide Sol：

有很多方法,詳細請看網路資料,值得思考  
<https://leetcode.com/problems/longest-common-prefix/solution/>

Approach 1: Horizontal scanning
1. 兩個字串先比較,取出共同prefix
2. 在用共同prefix與其他字串比較

Approach 3: Divide and conquer

## Reference
