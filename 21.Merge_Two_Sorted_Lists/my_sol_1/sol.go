package my_sol_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil {
		l1, l2 = l2, l1
	}

	result := l2

	for l1 != nil {
		for l2 != nil {
			if l2.Val > l1.Val {
				l2 = insertBefore(l2, l1.Val)
				result = l2
				break
			}

			//注意邊界條界,若下一個node為nil
			if l2.Next == nil && l1.Val >= l2.Val {
				insertAfter(l2, l1.Val)
				break
			}

			if l2.Next.Val >= l1.Val && l1.Val >= l2.Val {
				insertAfter(l2, l1.Val)
				break
			}
			l2 = l2.Next
		}
		l1 = l1.Next
	}
	return result
}

func insertAfter(node *ListNode, v int) {
	newNode := new(ListNode)
	newNode.Val = v
	newNode.Next = node.Next
	node.Next = newNode
}

func insertBefore(node *ListNode, v int) *ListNode {
	newNode := new(ListNode)
	newNode.Val = v
	newNode.Next = node
	return newNode
}
