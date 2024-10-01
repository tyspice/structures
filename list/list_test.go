package list

import (
	"slices"
	"testing"
)

func TestAddingAndReading(t *testing.T) {
	in := []string{"first", "second", "third"}
	out := make([]string, 0)

	testList := New(in[2])
	testList.AddNode(in[1])
	testList.AddNode(in[0])
	testList.ForEachNode(func(s string) {
		out = append(out, s)
	})

	if !slices.Equal(in, out) {
		t.Errorf(`FAIL: expected %v but got %v`, in, out)
	}
}
