package list

func newSentinal[T any]() *Element[T] {
	s := &Element[T]{}
	s.next = s
	s.prev = s
	return s
}
