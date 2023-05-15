package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	top, tail := head, head
	curr := 1
	node := head
	// var topPrev, tailPrev *ListNode
	for node != nil {
		// if curr == k-1 {
		// 	topPrev = node

		// }
		if curr == k {
			top = node
			tail = head

		}
		curr++
		fmt.Println(curr)
		if curr > k && node.Next != nil {
			// if tailPrev == nil {
			// 	tailPrev = head
			// } else {
			// 	tailPrev = tailPrev.Next
			// }
			tail = tail.Next
		}

		node = node.Next
	}
	tmp := top.Val
	// tmp := ListNode{Val: top.Val, Next: tail.Next}
	fmt.Println(top, tail)
	top.Val = tail.Val
	tail.Val = tmp
	// if topPrev == nil {
	// 	head = top
	// } else {
	// 	topPrev.Next = top
	// 	tailPrev.Next = &tmp
	// }
	return head
}

func main() {
	// n5 := ListNode{Val: 5, Next: nil}
	// n4 := ListNode{Val: 4, Next: &n5}
	// n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: nil}
	n1 := ListNode{Val: 1, Next: &n2}
	k := 1
	fmt.Println(swapNodes(&n1, k))
	// fmt.Println()
	node := &n1
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
