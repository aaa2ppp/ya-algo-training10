package main

import (
	"bytes"
	"io"
	"os"
	"unicode"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) int

func solve(a []int) int {
	n := len(a)
	_ = n
	return 0
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanIntLn(br, &n)

	var b bytes.Buffer
	for i := 0; i < n; i++ {
		b.Reset()
		s, _ := ScanString(br, '\n')
		for i, r := range s {
			if i == 0 {
				b.WriteRune(unicode.ToLower(r))
				continue
			}
			if unicode.IsUpper(r) {
				b.WriteByte('_')
				b.WriteRune(unicode.ToLower(r))
				continue
			}
			b.WriteRune(r)
		}
		b.WriteByte('\n')
		bw.Write(b.Bytes())
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
