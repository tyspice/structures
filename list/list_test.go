package list

import (
	"slices"
	"testing"
)

func TestAddingAndReading(t *testing.T) {
	in := []string{"first", "second", "third"}
	out := make([]string, 0)

	lst := New(in[2])
	lst.Insert(in[1])
	lst.Insert(in[0])
	lst.ForEach(func(s string) {
		out = append(out, s)
	})

	if !slices.Equal(in, out) {
		t.Errorf(`expected %v but got %v`, in, out)
	}

	if lst.head.data != in[0] {
		t.Errorf(`Head pointer was mutated. Expected %s but got %s`, in[0], lst.head.data)
	}
}
