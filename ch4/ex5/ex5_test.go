package ex5

import (
	"testing"
)

type TestData struct {
	data   []string
	result []string
}

func TestClearSameChar(t *testing.T) {
	tds := []TestData{
		{[]string{"I", "I", "D", "J", "J", "I", "I", "I", "G"}, []string{"I", "D", "J", "I", "G"}},
		{[]string{"I", "D", "D", "J", "D", "D", "D", "D", "G"}, []string{"I", "D", "J", "D", "G"}},
		{[]string{"I"}, []string{"I"}},
		{[]string{}, []string{}},
	}

	for k, td := range tds {
		if !SliceEqual(ClearSameChar(td.data), td.result) {
			t.Errorf("tds[%d]: want %v,but result is %v", k, td.result, ClearSameChar(td.data))
		}
	}
}

func SliceEqual(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for k := range x {
		if x[k] != y[k] {
			return false
		}
	}
	return true
}
