package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()

	fmt.Printf(" string %s,", b.String())
	r, _ := zlib.NewReader(&b)
	io.Copy(os.Stdout, r)
	r.Close()
}
