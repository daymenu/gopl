package main

import (
	"fmt"
	"sort"
)

type StringSclice []string

func (p StringSclice) Len() int {
	return len(p)
}

func (p StringSclice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p StringSclice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	names := []string{"z", "h", "a"}
	sort.Sort(StringSclice(names))
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)
}
