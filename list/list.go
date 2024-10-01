package list

type node[T any] struct {
	next *node[T]
	prev *node[T]
	data T
}

type LinkedList[T any] struct {
	head *node[T]
}

func New[T any](data T) LinkedList[T] {
	newNode := node[T]{data: data, next: nil, prev: nil}
	return LinkedList[T]{head: &newNode}
}

func (lst *LinkedList[T]) Insert(data T) {
	newNode := node[T]{next: lst.head, prev: nil, data: data}
	lst.head.prev = &newNode
	lst.head = &newNode
}

func (lst *LinkedList[T]) ForEach(fn func(T)) {
	currentNode := lst.head
	for currentNode != nil {
		fn(currentNode.data)
		currentNode = currentNode.next
	}
}
