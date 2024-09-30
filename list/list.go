package list

type node[T any] struct {
	next *node[T]
	data T
}

type LinkedList[T any] struct {
	head *node[T]
}

func New[T any](data T) LinkedList[T] {
	newNode := node[T]{data: data, next: nil}
	return LinkedList[T]{head: &newNode}
}

func (lst *LinkedList[T]) AddNode(data T) {
	newNode := node[T]{next: lst.head, data: data}
	lst.head = &newNode
}

func (lst *LinkedList[T]) ForEachNode(fn func(T)) {
	currentNode := lst.head
	for currentNode != nil {
		fn(currentNode.data)
		currentNode = currentNode.next
	}
}
