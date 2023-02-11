package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func reverseList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	newHead := reverseList(head.next)
	right := head.next
	right.next = head
	head.next = nil
	return newHead

}

func reverseInteratively(head *Node) *Node {
	var prev *Node
	prev = nil
	for head != nil {
		temp := head.next
		head.next = prev
		prev = head
		head = temp
	}
	return prev
}

func printList(head *Node) {
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}
func main() {
	arrayOfInt := []int{1, 2, 3, 4, 5}
	var head *Node
	var tail *Node
	for _, v := range arrayOfInt {
		nodeObj := Node{}
		nodeObj.val = v
		if head == nil {
			nodeObj.next = nil
			head = &nodeObj
		} else {
			nodeObj.next = nil
			tail.next = &nodeObj
		}
		tail = &nodeObj
	}
	printList(head)
	printList(reverseInteratively(head))
}
