package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func returnElemAtIndex(head *Node, index int) (*Node, error) {
	var targetNode *Node
	for index >= 0 {
		if head == nil {
			return nil, errors.New("Out of bounds!")
		}
		targetNode = head
		head = head.Next

		index--
	}

	return targetNode, nil

}
func main() {
	originalList := []int{1, 2, 3, 4, 5}
	var head *Node
	var tail *Node
	for _, v := range originalList {
		nodeObj := Node{}
		nodeObj.Value = v

		if head == nil {
			nodeObj.Next = nil
			nodeObj.Prev = nil
			head = &nodeObj
			tail = &nodeObj
		} else {
			nodeObj.Next = nil
			nodeObj.Prev = tail
			tail.Next = &nodeObj
			tail = &nodeObj
		}
	}
	targetNode, errorFetching := returnElemAtIndex(head, 1)
	if errorFetching != nil {
		fmt.Println(errorFetching)

	} else {
		fmt.Println(targetNode.Value)

	}
	fmt.Println("Forward")
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
	fmt.Println("Reverse")
	for tail != nil {
		fmt.Println(tail.Value)
		tail = tail.Prev
	}

}
