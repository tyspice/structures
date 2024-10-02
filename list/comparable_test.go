package list

import (
	"slices"
	"testing"
)

var (
	testComparableList     *ComparableList[string]
	testComparableListData = []string{"first", "second", "third"}
)

func TestNewComparableList(t *testing.T) {
	testComparableList = NewComparableList[string]()
	d := testComparableList.nil.Value
	if d != "" {
		t.Errorf(`expected %v but got %v`, "", d)
	}
}

func TestComparableList_Find(t *testing.T) {
	testComparableList.PushFront(testComparableListData[2])
	testComparableList.PushFront(testComparableListData[1])
	testComparableList.PushFront(testComparableListData[0])

	middle := testComparableList.Find(testComparableListData[1])
	if middle == nil {
		t.Fatal(`Result was nil`)
	}
	prevValue := middle.prev.Value
	nextValue := middle.next.Value
	if prevValue != testComparableListData[0] {
		t.Errorf(`expected prev value of %v but got %v`, testComparableListData[0], prevValue)
	}
	if nextValue != testComparableListData[2] {
		t.Errorf(`expected next value of %v but got %v`, testComparableListData[2], nextValue)
	}
}

func TestComparableList_FindAndDelete(t *testing.T) {
	err := testComparableList.FindAndDelete(testComparableListData[2])
	if err != nil {
		t.Fatal(`unable to find and delete value that should be findable`)
	}
	err = testComparableList.FindAndDelete(testComparableListData[2])
	if err == nil {
		t.Fatal(`found value that should have been deleted`)
	}
	expected := []string{testComparableListData[0], testComparableListData[1]}
	actual := []string{}
	testComparableList.ForEach(func(s string) {
		actual = append(actual, s)
	})
	if !slices.Equal(expected, actual) {
		t.Errorf(`expected %v but found %v`, expected, actual)
	}
}
