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

func hasCycle(head *ListNode) bool {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	// n5 := ListNode{Val: 5, Next: nil}
	// n4 := ListNode{Val: 4, Next: &n5}
	// n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: nil}
	n1 := ListNode{Val: 1, Next: &n2}
	n2.Next = &n1

	// node := hasCycle(&n1)
	// for node != nil {
	// 	fmt.Println(node.Val)
	// 	node = node.Next
	// }
	fmt.Println(hasCycle(&n1))
}
