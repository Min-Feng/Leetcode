package validateBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST_IncorrectSol(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 疏忽1. 沒考慮節點數值相同,此情況違反BST
	// 疏忽2. 右子樹所有節點一定大於跟節點,不能只單純看單一node的情況
	if (root.Left != nil) && (root.Left.Val >= root.Val) {
		return false
	}

	if (root.Right != nil) && (root.Right.Val <= root.Val) {
		return false
	}

	return isValidBST_IncorrectSol(root.Left) || isValidBST_IncorrectSol(root.Right)
}


func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if (root.Left != nil) && (root.Left.Val >= root.Val) {
		return false
	}

	if (root.Right != nil) && (root.Right.Val <= root.Val) {
		return false
	}

	return isValidBST(root.Left) || isValidBST(root.Right)
}