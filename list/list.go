package list

type N[T comparable] struct {
	next *N[T]
	prev *N[T]
	data T
}

type L[T comparable] struct {
	nil *N[T]
}

func New[T comparable]() *L[T] {
	sentinel := &N[T]{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &L[T]{nil: sentinel}
}

func (lst *L[T]) Insert(data T) {
	n := &N[T]{data: data}
	n.next = lst.nil.next
	n.prev = lst.nil
	lst.nil.next.prev = n
	lst.nil.next = n
}

func (lst *L[T]) ForEach(fn func(T)) {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		fn(x.data)
	}
}

func (lst *L[T]) Search(data T) (n *N[T]) {
	for x := lst.nil.next; x != lst.nil; x = x.next {
		if data == x.data {
			return x
		}
	}
	return nil
}
