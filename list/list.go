package list

type Element[T any] struct {
	next  *Element[T]
	prev  *Element[T]
	Value T
}

type List[T any] struct {
	nil *Element[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{nil: newSentinal[T]()}
}

func (lst *List[T]) PushFront(value T) *Element[T] {
	n := &Element[T]{Value: value}
	n.next = lst.nil.next // New node's next points to the current first node
	n.prev = lst.nil      // New node's prev points to the sentinel node
	lst.nil.next.prev = n // Current first node's prev points to the new node
	lst.nil.next = n      // Sentinel node's next points to the new node
	return n
}

func (lst *List[T]) ForEach(fn func(T)) {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		fn(x.Value)
	}
}

func (lst *List[T]) Delete(n *Element[T]) {
	n.prev.next = n.next
	n.next.prev = n.prev
}
