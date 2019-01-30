package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(Join(".", "han"))
}

func Join(sp string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	b := new(bytes.Buffer)
	b.WriteString(strs[0])
	for _, s := range strs[1:] {
		b.WriteString(sp + s)
	}

	return b.String()
}
