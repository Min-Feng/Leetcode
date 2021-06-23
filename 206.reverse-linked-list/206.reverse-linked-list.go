/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListV1(head *ListNode) *ListNode {
	var list []*ListNode
	var fresh *ListNode

	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	for i := len(list) - 1; i >= 0; i-- {
		if i == 0 {
			list[i].Next = nil
			fresh = list[len(list)-1]
			break
		}
		list[i].Next = list[i-1]
	}

	return fresh
}

func reverseListV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var current, tail, previous *ListNode

	current = head
	for current.Next != nil {
		previous = current
		current = current.Next // 當下就進入 next (current = current.Next), 變數命名感覺比較難理解
		previous.Next = tail
		tail = previous
	}

	current.Next = previous
	return current
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
// Memory Usage: 2.6 MB, less than 81.41% of Go online submissions for Reverse Linked List.
func reverseListV3(head *ListNode) *ListNode {
	var current, previous, next *ListNode

	current = head
	previous = nil
	for current != nil {
		next = current.Next // 重點技巧, 用一個 next 保存下一個節點的記憶體位置, 而 v2 的方式是 當下就進入 next (current = current.Next)
		current.Next = previous
		previous = current
		current = next // 做完所有的事情, 才前進到 下一個節點
	}

	return previous
}
