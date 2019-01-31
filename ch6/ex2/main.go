package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []int64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Len() int {
	l := 0
	for _, word := range s.words {
		for i := 0; i < 64; i++ {
			if word&(1<<uint(i)) != 0 {
				l++
			}
		}
	}
	return l
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, twrod := range t.words {
		if i < len(s.words) {
			s.words[i] |= twrod
		} else {
			s.words = append(s.words, twrod)
		}
	}
}

func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = append(t.words, s.words...)
	return &t
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) AddAll(ints ...int) {
	for _, y := range ints {
		s.Add(y)
	}
}
func main() {
	var s IntSet
	s.AddAll(1, 2, 3)
	fmt.Println(s.String())
}
