package list

import (
	"slices"
	"testing"
)

func testComparableListSetup() *ComparableList[string] {
	lst := NewComparableList[string]()
	d := []string{first, second, third}
	for _, e := range d {
		lst.PushBack(e)
	}
	return lst
}

func TestNewComparableList(t *testing.T) {
	lst := NewComparableList[string]()
	d := lst.nil.Value
	if d != "" {
		t.Errorf(`expected %v but got %v`, "", d)
	}
}

func TestComparableList_Find(t *testing.T) {
	lst := testComparableListSetup()
	lst.PushBack(first)
	lst.PushBack(second)
	lst.PushBack(third)

	middle := lst.Find(second)
	if middle == nil {
		t.Fatal(`Result was nil`)
	}
	prevValue := middle.prev.Value
	nextValue := middle.next.Value
	if prevValue != first {
		t.Errorf(`expected prev value of %v but got %v`, first, prevValue)
	}
	if nextValue != third {
		t.Errorf(`expected next value of %v but got %v`, third, nextValue)
	}
}

func TestComparableList_FindAndDelete(t *testing.T) {
	lst := testComparableListSetup()
	err := lst.FindAndDelete(third)
	if err != nil {
		t.Fatal(`unable to find and delete value that should be findable`)
	}
	err = lst.FindAndDelete(third)
	if err == nil {
		t.Fatal(`found value that should have been deleted`)
	}
	expected := []string{first, second}
	actual := []string{}
	lst.ForEach(func(s string) {
		actual = append(actual, s)
	})
	if !slices.Equal(expected, actual) {
		t.Errorf(`expected %v but found %v`, expected, actual)
	}
}
