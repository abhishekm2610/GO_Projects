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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var prev, curr, remove *ListNode
	nodes := 0
	curr = head
	for curr != nil {
		nodes++
		if nodes == n {
			remove = head
		}
		if nodes > n {
			prev = remove
			remove = prev.Next
		}
		curr = curr.Next
	}

	if remove == head {
		head = head.Next
	} else {
		prev.Next = remove.Next
	}
	return head
}

func main() {
	// n5 := ListNode{Val: 5, Next: nil}
	// n4 := ListNode{Val: 4, Next: &n5}
	// n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: nil}
	n1 := ListNode{Val: 1, Next: &n2}

	node := removeNthFromEnd(&n1, 1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
	// fmt.Println(reverseList(&n1))
}
