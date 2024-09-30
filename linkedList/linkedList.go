package linkedList

import "fmt"

type node struct {
	next *node
	data string
}

type linkedList struct {
	head *node
}

func NewLinkedList(data string) linkedList {
	newNode := node{data: data, next: nil}
	return linkedList{head: &newNode}
}

func (lst *linkedList) AddNode(data string) {
	newNode := node{next: lst.head, data: data}
	lst.head = &newNode
}

func (lst *linkedList) PrintNodeData() {
	currentNode := lst.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
