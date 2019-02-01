package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type limitreader struct {
	r io.Reader
	n int64
}

func (r *limitreader) Read(p []byte) (int, error) {
	r.r.Read(p[:r.n])
	return int(r.n), io.EOF
}
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitreader{r: r, n: n}
}
func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	rwc = w
	fmt.Printf("%T", w)
	// fmt.Println(html.Parse(LimitReader(strings.NewReader(`html`), 10)))
}
