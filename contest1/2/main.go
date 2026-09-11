package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(s, t string) int64

func solve(s, t string) int64 {
	m := make([]int, 256)

	for _, c := range t {
		m[c]++
	}

	var ans int64
	for r := 0; r < len(s); {
		c := s[r]
		if m[c] > 0 {
			r++
			m[c]--
			ans += int64(r)
			continue
		}
		if r > 0 {
			r--
			c = s[0]
			m[c]++
		}
		s = s[1:]
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	s, _ := ScanString(br, '\n')
	t, _ := ScanString(br, '\n')
	ans := solve(s, t)

	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
