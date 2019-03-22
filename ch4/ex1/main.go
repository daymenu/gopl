package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	xsha := sha256.Sum256([]byte("X"))
	xsum := 0
	for _, x := range xsha {
		for i := 0; i < 256; i++ {
			xsum += (int(x) & (1 << uint(i))) >> uint(i)
		}
	}
	fmt.Println(xsum)
}
