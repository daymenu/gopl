package mapequal

import "testing"

func TestEqual(t *testing.T) {
	x := map[string]int{
		"A": 0,
	}
	y := map[string]int{
		"B": 2,
	}
	if Equal(x, y) {
		t.Error(x, "!=", y)
	}
}
