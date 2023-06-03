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
	var prev *ListNode
	curr := head
	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode

	}
	return prev
}
func reorderList(head *ListNode) *ListNode {

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalf := slow.Next

	slow.Next = nil
	var prev *ListNode
	for secondHalf != nil {
		tmp := secondHalf.Next
		secondHalf.Next = prev
		prev = secondHalf
		secondHalf = tmp
	}

	first, second := head, prev

	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
	return head
}
func main() {
	n5 := ListNode{Val: 5, Next: nil}
	n4 := ListNode{Val: 4, Next: &n5}
	n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	node := reorderList(&n1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
	// fmt.Println(reverseList(&n1))
}
