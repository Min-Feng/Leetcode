package Approach_1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    MaxInt:=int(^uint(0) >> 1)
    MinInt:=^MaxInt
    
    return isValidBSTHelper(root,MaxInt,MinInt)
}

func isValidBSTHelper(root *TreeNode,max,min int) bool{
    if root == nil {
        return true
    }
    
    if (root.Val <= min || root.Val >= max){
        return false
    }
    
    return isValidBSTHelper(root.Left,root.Val,min) && isValidBSTHelper(root.Right,max,root.Val)
}