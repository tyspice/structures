package list

import (
	"slices"
	"testing"
)

var (
	testLst  *ComparableList[string]
	testData = []string{"first", "second", "third"}
)

func TestNewComparableList(t *testing.T) {
	testLst = NewComparableList[string]()
	d := testLst.nil.data
	if d != "" {
		t.Errorf(`expected %v but got %v`, "", d)
	}
}

func TestComparableList_Insert(t *testing.T) {
	third := testLst.Insert(testData[2])
	second := testLst.Insert(testData[1])
	head := testLst.Insert(testData[0])
	nodes := []string{head.data, second.data, third.data}
	for i := range testData {
		if nodes[i] != testData[i] {
			t.Errorf(`expected %v but got %v`, testData[1], nodes[i])
		}
	}
}

func TestComparableList_ForEach(t *testing.T) {
	out := make([]string, 0)

	testLst.ForEach(func(s string) {
		out = append(out, s)
	})

	if !slices.Equal(testData, out) {
		t.Errorf(`expected %v but got %v`, testData, out)
	}

	if testLst.nil.next.data != testData[0] {
		t.Errorf(`Head pointer was mutated. Expected %s but got %s`, testData[0], testLst.nil.next.data)
	}
}

func TestComparableList_Find(t *testing.T) {
	middle := testLst.Find(testData[1])
	if middle == nil {
		t.Fatal(`Result was nil`)
	}
	prevData := middle.prev.data
	nextData := middle.next.data
	if prevData != testData[0] {
		t.Errorf(`expected prev value of %v but got %v`, testData[0], prevData)
	}
	if nextData != testData[2] {
		t.Errorf(`expected next value of %v but got %v`, testData[2], nextData)
	}
}

func TestComparableList_Delete(t *testing.T) {
	middle := testLst.Find(testData[1])
	testLst.Delete(middle)
	expected := []string{testData[0], testData[2]}
	actual := []string{}
	testLst.ForEach(func(d string) {
		actual = append(actual, d)
	})
	if !slices.Equal(expected, actual) {
		t.Errorf(`expected %v but got %v`, expected, actual)
	}
}

func TestComparableList_FindAndDelete(t *testing.T) {
	err := testLst.FindAndDelete(testData[2])
	if err != nil {
		t.Fatal(`unable to find and delete value that should be findable`)
	}
	err = testLst.FindAndDelete(testData[2])
	if err == nil {
		t.Fatal(`found value that should have been deleted`)
	}
	expected := []string{testData[0]}
	actual := []string{}
	testLst.ForEach(func(s string) {
		actual = append(actual, s)
	})
	if !slices.Equal(expected, actual) {
		t.Errorf(`expected %v but found %v`, expected, actual)
	}
}
