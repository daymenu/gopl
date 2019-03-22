package rev

import (
	"testing"
)

type Rev struct {
	s    []int
	want []int
}

func TestReverse(t *testing.T) {
	silces := getSlice()
	for _, s := range silces {
		Reverse(s.s)
		for i := 0; i < len(s.s); i++ {
			if s.s[i] != s.want[i] {
				t.Errorf("%v\t%v", s.s, s.want)
			}
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		silces := getSlice()
		for _, s := range silces {
			Reverse(s.s)
			for i := 0; i < len(s.s); i++ {
				if s.s[i] != s.want[i] {
					b.Errorf("%v\t%v", s.s, s.want)
				}
			}
		}
	}
}

func getSlice() []Rev {
	return []Rev{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{nil, nil},
	}
}
