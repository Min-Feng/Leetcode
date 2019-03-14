package Approach_2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var inOrderList []int
	var stack []*TreeNode
	isTraversal := make(map[*TreeNode]bool)
	stack = append(stack, root)

	for len(stack) != 0 {
		for root.Left != nil && isTraversal[root.Left] == false {
			stack = append(stack, root.Left)
			root = root.Left
			isTraversal[root] = true
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		inOrderList = append(inOrderList, root.Val)

		if n := len(inOrderList); n > 1 {
			preValue := inOrderList[n-2]
			nowValue := inOrderList[n-1]
			if preValue >= nowValue {
				return false
			}
		}

		if root.Right != nil && isTraversal[root.Right] == false {
			stack = append(stack, root.Right)
			root = root.Right
			isTraversal[root] = true
		}

	}

	return true
}

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var values []int
	values = inOrder(root, values)
	var pre int
	for i, v := range values {
		if i == 0 {
			pre = v
			continue
		}
		if pre >= v {
			return false
		}
		pre = v
	}
	return true
}
func inOrder(root *TreeNode, values []int) []int {
	if root != nil {
		values = inOrder(root.Left, values)
		values = append(values, root.Val)
		values = inOrder(root.Right, values)
	}
	return values
}
