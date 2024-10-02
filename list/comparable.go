package list

import "errors"

type CNode[T comparable] struct {
	next *CNode[T]
	prev *CNode[T]
	Data T
}

type Comparable[T comparable] struct {
	nil *CNode[T]
}

func NewComparable[T comparable]() *Comparable[T] {
	sentinel := &CNode[T]{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &Comparable[T]{nil: sentinel}
}

func (lst *Comparable[T]) PushFront(data T) *CNode[T] {
	n := &CNode[T]{Data: data}
	n.next = lst.nil.next // New node's next points to the current first node
	n.prev = lst.nil      // New node's prev points to the sentinel node
	lst.nil.next.prev = n // Current first node's prev points to the new node
	lst.nil.next = n      // Sentinel node's next points to the new node
	return n
}

func (lst *Comparable[T]) ForEach(fn func(T)) {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		fn(x.Data)
	}
}

func (lst *Comparable[T]) Find(data T) *CNode[T] {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		if data == x.Data {
			return x
		}
	}
	return nil
}

func (lst *Comparable[T]) Delete(n *CNode[T]) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lst *Comparable[T]) FindAndDelete(data T) error {
	n := lst.Find(data)
	if n == nil {
		return errors.New("item not found")
	}
	lst.Delete(n)
	return nil
}
