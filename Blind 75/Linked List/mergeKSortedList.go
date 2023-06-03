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

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	for len(lists) > 1 {
		merged := mergeTwoLists(lists[0], lists[1])
		lists = lists[2:]
		lists = append(lists, merged)
	}
	return lists[0]

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
	n2 := ListNode{Val: 3, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	n23 := ListNode{Val: 5, Next: nil}
	n22 := ListNode{Val: 4, Next: &n23}
	n21 := ListNode{Val: 1, Next: &n22}

	n32 := ListNode{Val: 6, Next: nil}
	n31 := ListNode{Val: 2, Next: &n32}

	lists := []*ListNode{&n1, &n21, &n31}
	node := mergeKLists(lists)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
	// fmt.Println(hasCycle(&n1))
}
