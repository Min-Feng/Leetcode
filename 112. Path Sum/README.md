# 112. Path Sum

## 原題連結
[https://leetcode.com/problems/path-sum/description/](https://leetcode.com/problems/path-sum/description/)

## My Solution
實做成功,方式如同Approach_1  
Approach_2,參考解答進行練習

## Approach_1：[Link](Approach_1/main.go)  
針對每個node進行條件判斷,使用遞迴(Recursion)進行DFS,個人想到的解法

## Approach_2：[Link](Approach_2/main.go)   
利用stack,以Iterations,進行DFS  
hasPathSum 只使用一個stack其元素的型別為空界面 interface{}  
hasPathSum2 使用兩個stack,個別維護

*實做錯誤：*
1. 短變數宣告的誤解,是否有block的存在,會有不同效果,細節請點 [Ref](https://play.golang.org/p/Z76cc5QJFbx)  
不應該在迴圈內宣告一個內部的stack,應該在外部就先宣告好,  
迴圈判斷是參考外部變數,內部變數無法影響外部  

