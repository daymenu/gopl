package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	bits := []struct {
		num  uint64
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
	}
	for _, v := range bits {
		if PopCount(v.num) != v.want {
			t.Errorf("PopCount(%b)=%d,want %d", v.num, PopCount(v.num), v.want)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {

}
