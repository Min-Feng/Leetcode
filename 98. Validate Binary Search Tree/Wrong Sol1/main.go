package wrong_sol1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 疏忽1. 沒考慮節點數值相同,此情況違反BST
	// 疏忽2. 不能只單純看單一node與父子節點的大小關係
	if (root.Left != nil) && (root.Left.Val >= root.Val) {
		return false
	}

	if (root.Right != nil) && (root.Right.Val <= root.Val) {
		return false
	}

	return isValidBST(root.Left) || isValidBST(root.Right)
}