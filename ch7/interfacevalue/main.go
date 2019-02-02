package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//接口值有动态类型和动态值
func main() {
	var w io.Writer
	fmt.Printf("%T,%[1]v\n", w)
	w = os.Stdout
	fmt.Printf("%T,%[1]q\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("%T,%[1]q\n", w)
	w = nil
	fmt.Printf("%T,%[1]v\n", w)
	const debug = true
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		fmt.Println(buf.String())
	}
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

// 理解，接口值分为两部分接口类型和接口值，只有两个都为nil才为nil
// 普通类型变量，一定会有类型，所以值为nil就是nil
