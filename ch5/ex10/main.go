package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for k, _ := range items {
			if !seen[k] {
				seen[k] = true
				mm := make(map[string]bool, len(m[k]))
				for _, kk := range m[k] {
					mm[kk] = true
				}
				visitAll(mm)
				order = append(order, k)
			}
		}
	}

	keys := make(map[string]bool)
	for k, _ := range m {
		keys[k] = true
	}
	visitAll(keys)
	return order
}
