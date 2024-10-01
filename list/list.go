package list

import "errors"

type ComparableNode[T comparable] struct {
	next *ComparableNode[T]
	prev *ComparableNode[T]
	data T
}

type ComparableList[T comparable] struct {
	nil *ComparableNode[T]
}

func New[T comparable]() *ComparableList[T] {
	sentinel := &ComparableNode[T]{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &ComparableList[T]{nil: sentinel}
}

func (lst *ComparableList[T]) Insert(data T) *ComparableNode[T] {
	n := &ComparableNode[T]{data: data}
	n.next = lst.nil.next
	n.prev = lst.nil
	lst.nil.next.prev = n
	lst.nil.next = n
	return n
}

func (lst *ComparableList[T]) ForEach(fn func(T)) {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		fn(x.data)
	}
}

func (lst *ComparableList[T]) Find(data T) *ComparableNode[T] {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		if data == x.data {
			return x
		}
	}
	return nil
}

func (lst *ComparableList[T]) Delete(n *ComparableNode[T]) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lst *ComparableList[T]) FindAndDelete(data T) error {
	n := lst.Find(data)
	if n == nil {
		return errors.New("item not found")
	}
	lst.Delete(n)
	return nil
}
