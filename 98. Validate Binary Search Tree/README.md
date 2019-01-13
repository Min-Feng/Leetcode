# 98. Validate Binary Search Tree

## 原題連結
[https://leetcode.com/problems/validate-binary-search-tree/description/](https://leetcode.com/problems/validate-binary-search-tree/description/)

## My Solution
實做失敗

*實做錯誤：*
- **Wrong Sol 1：[Link](Wrong%20Sol1/main.go)**
1. 沒考慮節點數值相同,此情況違反BST
2. 右子樹所有節點一定大於根節點,不能只單純看單一node與父子節點的大小關係

- **Wrong Sol 2：[Link](Wrong%20Sol2/main.go)**
3. 藉由傳入父節點,比較了父節點跟子節點的大小,但應該考慮整體,右子樹所有節點一定大於左邊所有節點

## Approach_1：[Link](Approach_1/main.go)  
應該新設一個輔助函數來傳入每個節點該遵守的最大,最小值範圍,例如  
根節點其數值應該介於,MinInt< root.Val < MaxInt  
無符號整數的最小值0,進行位元反轉^,得到最大值,又因為有號整數第一個bit代表正負號,故右移動1bit
````go
MaxInt:=int(^uint(0) >> 1)  
MinInt:=^MaxInt
````
依此類推  
根節點左子樹其數值應該介於,MinInt< root.Left.Val < root.Val  
根節點右子樹其數值應該介於,root.Val< root.Right.Val < MaxInt  

## Approach_2：[Link](Approach_2/main.go)   
由於BST中序走訪後，一定是排序過大小的數列  
進行中序走訪,若走訪過程檢查到沒有進行大小排序,則可認定不屬於BST  
效能較差,原因沒想到為什麼

## Reference:
[https://shannonhu.gitbooks.io/go-leetcode/content/98_validate_a_binary_search_tree.html](https://shannonhu.gitbooks.io/go-leetcode/content/98_validate_a_binary_search_tree.html)

