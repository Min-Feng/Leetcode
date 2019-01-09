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

	var item []interface{}
	var curNode *TreeNode
	var curSum int

	for len(stack) != 0 {
		item, stack = stack[len(stack)-1], stack[:len(stack)-1]
		curNode = item[0].(*TreeNode)
		curSum = item[1].(int)

		if curSum == 0 && curNode.Left == nil && curNode.Right == nil {
			return true
		}

		if curNode.Left != nil {
			stack = append(stack, []interface{}{curNode.Left, curSum - curNode.Left.Val})
		}

		if curNode.Right != nil {
			stack = append(stack, []interface{}{curNode.Right, curSum - curNode.Right.Val})
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

	var cur_node *TreeNode
	var cur_sum int

	for len(nodeStack) != 0 {
		cur_node, nodeStack = nodeStack[len(nodeStack)-1], nodeStack[:len(nodeStack)-1]
		cur_sum, sumStack = sumStack[len(sumStack)-1], sumStack[:len(sumStack)-1]

		if cur_sum == 0 && cur_node.Left == nil && cur_node.Right == nil {
			return true
		}

		if cur_node.Left != nil {
			nodeStack = append(nodeStack, cur_node.Left)
			sumStack = append(sumStack, cur_sum-cur_node.Left.Val)
		}

		if cur_node.Right != nil {
			nodeStack = append(nodeStack, cur_node.Right)
			sumStack = append(sumStack, cur_sum-cur_node.Right.Val)
		}
	}
	return false
}
