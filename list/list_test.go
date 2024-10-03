package list

import (
	"slices"
	"testing"
)

var (
	testList         *List[string]
	testListData     = []string{"first", "second", "third"}
	headElementRef   *Element[string]
	MiddleElementRef *Element[string]
	lastElementRef   *Element[string]
)

func TestNewList(t *testing.T) {
	testList = NewList[string]()
	d := testList.nil.Value
	if d != "" {
		t.Errorf(`expected %v but got %v`, "", d)
	}
}

func TestList_PushFront(t *testing.T) {
	lastElementRef = testList.PushFront(testComparableListData[2])
	MiddleElementRef = testList.PushFront(testComparableListData[1])
	headElementRef = testList.PushFront(testComparableListData[0])
	nodes := []string{headElementRef.Value, MiddleElementRef.Value, lastElementRef.Value}
	for i := range testListData {
		if nodes[i] != testListData[i] {
			t.Errorf(`expected %v but got %v`, testComparableListData[1], nodes[i])
		}
	}
}

func TestList_ForEach(t *testing.T) {
	out := []string{}

	testList.ForEach(func(s string) {
		out = append(out, s)
	})

	if !slices.Equal(testListData, out) {
		t.Errorf(`expected %v but got %v`, testListData, out)
	}

	if testList.nil.next.Value != testListData[0] {
		t.Errorf(`Head pointer was mutated. Expected %s but got %s`, testListData[0], testList.nil.next.Value)
	}
}

func TestList_Delete(t *testing.T) {
	testList.Delete(MiddleElementRef)
	expected := []string{testListData[0], testListData[2]}
	actual := []string{}
	testList.ForEach(func(d string) {
		actual = append(actual, d)
	})
	if !slices.Equal(expected, actual) {
		t.Errorf(`expected %v but got %v`, expected, actual)
	}
}

func TestList_PushBack(t *testing.T) {
	testList.PushBack(testListData[1])
	expected := []string{testListData[0], testListData[2], testListData[1]}
	actual := []string{}
	testList.ForEach(func(d string) {
		actual = append(actual, d)
	})
	if !slices.Equal(expected, actual) {
		t.Errorf(`expected %v but got %v`, expected, actual)
	}
}
