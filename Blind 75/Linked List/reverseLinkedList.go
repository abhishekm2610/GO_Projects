/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// prev.Next = head

	var recursive func(prev *ListNode, curr *ListNode) *ListNode

	recursive = func(prev *ListNode, curr *ListNode) *ListNode {
		if curr == nil {
			return prev
		}
		tmp := curr.Next
		curr.Next = prev
		return recursive(curr, tmp)
	}

	return recursive(nil, head)

	// curr := head
	// for curr != nil {
	// 	nextNode := curr.Next
	// 	curr.Next = prev
	// 	prev = curr
	// 	curr = nextNode

	// }

	// for head != nil {
	// 	newNode := ListNode{Val: head.Val, Next: curr}
	// 	curr = &newNode
	// 	head = head.Next
	// }
}

func main() {
	n5 := ListNode{Val: 5, Next: nil}
	n4 := ListNode{Val: 4, Next: &n5}
	n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	node := reverseList(&n1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
	// fmt.Println(reverseList(&n1))
}
