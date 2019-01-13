# LeetCodeTools
建立leetCode常用到的資料結構,以及相關資料結構常用到的方法

## type TreeNode
二元樹結構
- **func List2TreeNode(values ...interface{}) \*TreeNode**  
  利用階層尋訪的步驟,插入數值到二元樹  
  insertLevelOrder()類似BFS,所以應該使用queue來幫助程式撰寫  
  自己實做時,採用了遞迴的方式,也就是使用了DFS的方式,所以思考方式變得不直觀  
  以後實做要先想清楚,題目需求是以什麼樣的方式尋訪！！！
  [Reference](http://alrightchiu.github.io/SecondRound/binary-tree-traversalxun-fang.html#level)
