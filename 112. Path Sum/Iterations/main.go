package iterations

// https://leetcode.com/problems/path-sum/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// with the help of stack. DFS would be better than BFS
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	var stack [][]interface{}
	sum -= root.Val
	stack = append(stack, []interface{}{root, sum})

	for len(stack) != 0 {
		item, stack := stack[len(stack)-1], stack[:len(stack)-1]
		node := item[0].(*TreeNode)
		sum = item[1].(int)

		if sum == 0 && node.Left == nil && node.Right == nil {
			return true
		}

		if node.Left != nil {
			stack = append(stack, []interface{}{node.Left, sum - node.Left.Val})
		}

		if node.Right != nil {
			stack = append(stack, []interface{}{node.Right, sum - node.Right.Val})
		}
	}
	return false
}

func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	var nodeStack []*TreeNode
	var sumStack []int

	nodeStack = append(nodeStack, root)
	sumStack = append(sumStack, sum-root.Val)

	for len(nodeStack) != 0 {
		node, nodeStack := nodeStack[len(nodeStack)-1], nodeStack[:len(nodeStack)-1]
		sum, sumStack = sumStack[len(sumStack)-1], sumStack[:len(sumStack)-1]

		if sum == 0 && node.Left == nil && node.Right == nil {
			return true
		}

		if node.Left != nil {
			nodeStack = append(nodeStack, root.Left)
			sumStack = append(sumStack, sum-root.Left.Val)
		}

		if node.Right != nil {
			nodeStack = append(nodeStack, root.Right)
			sumStack = append(sumStack, sum-root.Right.Val)
		}
	}
	return false
}
