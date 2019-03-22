package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	fun := flag.Int("f", 1, "1: sha256 \t2:sha384 \t3:sha:512")
	s := flag.String("s", "", "请输入要加密的字符串")
	flag.Parse()
	switch *fun {
	case 1:
		fmt.Printf("%x\n", sha256.Sum256([]byte(*s)))
	case 2:
		fmt.Printf("%x\n", sha512.Sum384([]byte(*s)))
	case 3:
		fmt.Printf("%x\n", sha512.Sum512([]byte(*s)))
	}
}
