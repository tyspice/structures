package list

import "errors"

type ComparableList[T comparable] struct {
	List[T]
}

func NewComparableList[T comparable]() *ComparableList[T] {
	return &ComparableList[T]{List[T]{nil: newSentinal[T]()}}
}

func (lst *ComparableList[T]) Find(value T) *Element[T] {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		if value == x.Value {
			return x
		}
	}
	return nil
}

func (lst *ComparableList[T]) FindAndDelete(value T) error {
	n := lst.Find(value)
	if n == nil {
		return errors.New("item not found")
	}
	lst.Delete(n)
	return nil
}
