package list

type Element[T any] struct {
	next  *Element[T]
	prev  *Element[T]
	Value T
}

func (e *Element[T]) Next() *Element[T] {
	return e.next
}

func (e *Element[T]) Prev() *Element[T] {
	return e.prev
}

type List[T any] struct {
	nil *Element[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{nil: newSentinal[T]()}
}

func (lst *List[T]) Front() *Element[T] {
	if lst.nil.next == lst.nil {
		return nil
	}
	return lst.nil.next
}

func (lst *List[T]) Back() *Element[T] {
	if lst.nil.prev == lst.nil {
		return nil
	}
	return lst.nil.prev
}

func (lst *List[T]) PushFront(value T) *Element[T] {
	n := &Element[T]{Value: value}
	n.next = lst.nil.next // New node's next points to the current first node
	n.prev = lst.nil      // New node's prev points to the sentinel node
	lst.nil.next.prev = n // Current first node's prev points to the new node
	lst.nil.next = n      // Sentinel node's next points to the new node
	return n
}

func (lst *List[T]) PushBack(value T) *Element[T] {
	n := &Element[T]{Value: value}
	n.next = lst.nil      // New node's next points to the sentinel
	n.prev = lst.nil.prev // New node's prev points to the current last node
	lst.nil.prev.next = n // Current last node's next points to the new node
	lst.nil.prev = n      // Sentinel node's prev points to the new node
	return n
}

func (lst *List[T]) ForEach(fn func(T)) {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		fn(x.Value)
	}
}

func (lst *List[T]) Remove(n *Element[T]) {
	n.prev.next = n.next
	n.next.prev = n.prev
}
