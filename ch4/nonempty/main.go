package main

import "fmt"

func main() {
	strs := []string{"H", "a", "n", "", "J"}
	fmt.Println(nonempty(strs))
	fmt.Println(strs)
}

func nonempty(strs []string) []string {

	i := 0
	for _, s := range strs {
		if s != "" {
			strs[i] = s
			i++
		}
	}
	return strs[:i]
}
