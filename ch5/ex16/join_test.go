package main

import (
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join("/", "han", "jian", "han", "jian", "han", "jian", "han", "jian", "han", "jian", "han", "jian")
		// strings.Join([]string{"han", "jian", "han", "jian", "han", "jian", "han", "jian", "han", "jian", "han", "jian"}, "/")
	}
}
