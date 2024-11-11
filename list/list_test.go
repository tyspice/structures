package list

import (
	"slices"
	"testing"
)

const (
	first  = "first"
	second = "second"
	third  = "third"
)

func testListSetup(d ...string) *List[string] {
	if len(d) == 0 {
		d = []string{first, second, third}
	}
	lst := NewList[string]()
	for _, e := range d {
		lst.PushBack(e)
	}
	return lst
}

func TestNewList(t *testing.T) {
	lst := NewList[string]()
	d := lst.nil.Value
	if d != "" {
		t.Errorf(`expected %v but got %v`, "", d)
	}
}

func TestList_PushBack(t *testing.T) {
	expected := []string{first, third, second}
	lst := testListSetup(first, third, second)
	first := lst.nil.next
	second := first.next
	third := second.next
	actual := []string{first.Value, second.Value, third.Value}
	if !slices.Equal(expected, actual) {
		t.Errorf(`expected %v but got %v`, expected, actual)
	}
}

func TestList_PushFront(t *testing.T) {
	lst := NewList[string]()
	lastElementRef := lst.PushFront(third)
	MiddleElementRef := lst.PushFront(second)
	headElementRef := lst.PushFront(first)
	nodes := []string{headElementRef.Value, MiddleElementRef.Value, lastElementRef.Value}
	data := []string{first, second, third}
	for i := range data {
		if nodes[i] != data[i] {
			t.Errorf(`expected %v but got %v`, data[i], nodes[i])
		}
	}
}

func TestList_ForEach(t *testing.T) {
	out := []string{}
	expected := []string{first, second, third}
	lst := testListSetup()
	lst.ForEach(func(s string) {
		out = append(out, s)
	})

	if !slices.Equal(expected, out) {
		t.Errorf(`expected %v but got %v`, expected, out)
	}

	if lst.nil.next.Value != expected[0] {
		t.Errorf(`Head pointer was mutated. Expected %s but got %s`, expected[0], lst.nil.next.Value)
	}
}

func TestList_Remove(t *testing.T) {
	lst := testListSetup()
	middle := lst.nil.next.next
	lst.Remove(middle)
	expected := []string{first, third}
	actual := []string{}
	lst.ForEach(func(d string) {
		actual = append(actual, d)
	})
	if !slices.Equal(expected, actual) {
		t.Errorf(`expected %v but got %v`, expected, actual)
	}
}

func TestList_Front(t *testing.T) {
	lst := testListSetup()
	value := lst.Front().Value
	if value != first {
		t.Errorf(`expected %v but got %v`, first, value)
	}
}

func TestList_Back(t *testing.T) {
	lst := testListSetup()
	value := lst.Back().Value
	if value != third {
		t.Errorf(`expected %v but got %v`, third, value)
	}
}
