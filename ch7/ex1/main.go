package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		*w++
	}
	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		*l++
	}
	return len(p), nil
}
func main() {
	var w WordCounter
	fmt.Fprintf(&w, "i am a boy")
	fmt.Println(w)
	var l LineCounter
	fmt.Fprintf(&l, `i am a boy!
	are you a girl?`)
	fmt.Println(l)
}
