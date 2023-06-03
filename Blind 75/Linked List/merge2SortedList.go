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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newNode ListNode
	var head *ListNode
	head = &newNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}

	if list1 == nil && list2 != nil {
		head.Next = list2
	} else if list2 == nil && list1 != nil {
		head.Next = list1
	}

	return newNode.Next
}

func main() {
	// n5 := ListNode{Val: 5, Next: nil}
	// n4 := ListNode{Val: 4, Next: &n5}
	n3 := ListNode{Val: 4, Next: nil}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	n23 := ListNode{Val: 4, Next: nil}
	n22 := ListNode{Val: 3, Next: &n23}
	n21 := ListNode{Val: 1, Next: &n22}

	node := mergeTwoLists(&n1, &n21)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
	// fmt.Println(hasCycle(&n1))
}
