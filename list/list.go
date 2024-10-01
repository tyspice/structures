package list

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	data T
}

type LinkedList[T any] struct {
	nil *Node[T]
}

func New[T any]() *LinkedList[T] {
	sentinel := &Node[T]{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &LinkedList[T]{nil: sentinel}
}

func (lst *LinkedList[T]) Insert(data T) {
	n := &Node[T]{data: data}
	n.next = lst.nil.next
	n.prev = lst.nil
	lst.nil.next.prev = n
	lst.nil.next = n
}

func (lst *LinkedList[T]) ForEach(fn func(T)) {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		fn(x.data)
	}
}
