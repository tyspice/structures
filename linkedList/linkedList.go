package linkedList

import "fmt"

type node struct {
	next *node
	data string
}

type LinkedList struct {
	head *node
}

func NewLinkedList(data string) LinkedList {
	newNode := node{data: data, next: nil}
	return LinkedList{head: &newNode}
}

func (lst *LinkedList) AddNode(data string) {
	newNode := node{next: lst.head, data: data}
	lst.head = &newNode
}

func (lst *LinkedList) PrintNodeData() {
	currentNode := lst.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
