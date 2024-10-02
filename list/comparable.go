package list

import "errors"

type CNode[T comparable] struct {
	next *CNode[T]
	prev *CNode[T]
	data T
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

func (lst *Comparable[T]) Insert(data T) *CNode[T] {
	n := &CNode[T]{data: data}
	n.next = lst.nil.next
	n.prev = lst.nil
	lst.nil.next.prev = n
	lst.nil.next = n
	return n
}

func (lst *Comparable[T]) ForEach(fn func(T)) {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		fn(x.data)
	}
}

func (lst *Comparable[T]) Find(data T) *CNode[T] {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		if data == x.data {
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
