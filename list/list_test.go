package list

import (
	"slices"
	"testing"
)

var (
	testLst  LinkedList[string]
	testData = []string{"first", "second", "third"}
)

func TestNew(t *testing.T) {
	testLst = New(testData[2])
	d := testLst.head.data
	if d != testData[2] {
		t.Errorf(`expected %v but got %v`, testData[2], d)
	}
}

func TestInsert(t *testing.T) {
	testLst.Insert(testData[1])
	testLst.Insert(testData[0])
	head := testLst.head
	second := testLst.head.next
	third := second.next
	nodes := []string{head.data, second.data, third.data}
	for i := range testData {
		if nodes[i] != testData[i] {
			t.Errorf(`expected %v but got %v`, testData[1], nodes[i])
		}
	}
}

func TestForEach(t *testing.T) {
	out := make([]string, 0)

	testLst.ForEach(func(s string) {
		out = append(out, s)
	})

	if !slices.Equal(testData, out) {
		t.Errorf(`expected %v but got %v`, testData, out)
	}

	if testLst.head.data != testData[0] {
		t.Errorf(`Head pointer was mutated. Expected %s but got %s`, testData[0], testLst.head.data)
	}
}
