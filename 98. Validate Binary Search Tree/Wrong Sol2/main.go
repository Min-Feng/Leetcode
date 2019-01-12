package wrong_sol2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTPassParent(root, nil)
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