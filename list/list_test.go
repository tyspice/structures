package list

import (
	"slices"
	"testing"
)

var (
	testLst  *L[string]
	testData = []string{"first", "second", "third"}
)

func TestNew(t *testing.T) {
	testLst = New[string]()
	d := testLst.nil.data
	if d != "" {
		t.Errorf(`expected %v but got %v`, "", d)
	}
}

func TestL_Insert(t *testing.T) {
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

func TestL_ForEach(t *testing.T) {
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

func TestL_Search(t *testing.T) {
	middle := testLst.Search(testData[1])
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
