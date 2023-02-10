package main

import "fmt"

type Node struct {
	val  int
	next *Node
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
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}
