package ex3

import (
	"testing"
)

func TestReverse(t *testing.T) {
	arr := [RLen]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Reverse(&arr)
	t.Error(arr)
}
