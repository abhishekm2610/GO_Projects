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

func printList(head *Node) {
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}
func main() {
	arrayOfInt := []int{}
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
	printList(reverseList(head))
}
