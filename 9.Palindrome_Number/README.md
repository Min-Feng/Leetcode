# 9. Palindrome Number

## 原題連結

<https://leetcode.com/problems/palindrome-number/>

## My Solution

my_Sol_1：[Link](my_sol_1/my_sol_1.go)

1. 利用除法得商數跟餘數
2. 將每次得到的餘數放入到slice中
3. for迴圈循環直到商數為0
4. 比較slice內部是否符合條件
5. 使用的記憶體空間較多

## Website Provide Sol：

1. 利用除法得商數跟餘數
2. 把餘數當成newValue,每次迴圈重新進行計算
3. newValue=newValue*10+remainder
4. 當newValue > 更新後的x number,迴圈停止
5. 使用的記憶體空間較少

## Reference

練習使用測試,記得使用-v才可以顯示完整訊息  
目錄路徑有空白的話,會造成go test指令失敗  
<https://blog.wu-boy.com/2018/05/how-to-write-testing-in-golang/>