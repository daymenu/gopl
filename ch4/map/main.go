package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"liming":   23,
		"gutianle": 43,
	}
	// fmt.Println(ages)

	// delete(ages, "liming")
	// fmt.Println(ages)

	names := make([]string, 0, len(ages))

	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println("--", ages[name])
	}

	var stus map[string]int
	fmt.Println(stus == nil)
	delete(stus, "dd")
}
