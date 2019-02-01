package main

import (
	"fmt"
	"io"
	"os"
)

type Wcounter struct {
	w       io.Writer
	counter int64
}

func (w *Wcounter) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.counter += int64(n)
	return n, err
}
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wr := &Wcounter{w: w}
	return wr, &wr.counter
}

func main() {
	w, c := CountingWriter(os.Stdout)
	w.Write([]byte("hanjian\n"))
	fmt.Println(*c)
}
