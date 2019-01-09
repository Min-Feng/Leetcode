package validateBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTPassParent(root, nil)
}

func isValidBST_IncorrectSol(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 疏忽1. 沒考慮節點數值相同,此情況違反BST
	// 疏忽2. 右子樹所有節點一定大於根節點,不能只單純看單一node的情況
	if (root.Left != nil) && (root.Left.Val >= root.Val) {
		return false
	}

	if (root.Right != nil) && (root.Right.Val <= root.Val) {
		return false
	}

	return isValidBST_IncorrectSol(root.Left) || isValidBST_IncorrectSol(root.Right)
}

func isValidBSTPassParent(root, parent *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		// 疏忽3. 右子樹所有節點一定大於左邊所有節點
		if parent != nil && root.Left.Val >= parent.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		if parent != nil && root.Right.Val <= parent.Val {
			return false
		}
	}

	return isValidBSTPassParent(root.Left, root) || isValidBSTPassParent(root.Right, root)
}
