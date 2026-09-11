package main

import (
	"io"
	"os"
	"strings"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(s string) (string, bool)

func solve(s string) (string, bool) {
	lets := make([]int, 26)
	var total, open, close, slash int

	for _, c := range []byte(s) {
		switch c {
		case '<':
			open++
		case '>':
			close++
		case '/':
			slash++
		default:
			total++
			lets[c-'a']++
		}
	}

	if open == 0 {
		return "", false
	}
	if open%2 != 0 {
		return "", false
	}
	if close != open {
		return "", false
	}
	if slash*2 != open {
		return "", false
	}
	if total < open {
		return "", false
	}
	for i := range lets {
		if lets[i]%2 != 0 {
			return "", false
		}
	}

	var b strings.Builder
	k := 0
	for i := 2; i < open; i += 2 {
		for lets[k] == 0 {
			k++
		}
		c := byte('a' + k)
		lets[k] -= 2
		b.WriteByte('<')
		b.WriteByte(c)
		b.WriteByte('>')
		b.WriteString("</")
		b.WriteByte(c)
		b.WriteByte('>')
	}

	var last strings.Builder
	for k < 26 {
		for i := 0; i < lets[k]; i += 2 {
			last.WriteByte(byte('a' + k))
		}
		k++
	}

	b.WriteByte('<')
	b.WriteString(last.String())
	b.WriteByte('>')
	b.WriteString("</")
	b.WriteString(last.String())
	b.WriteByte('>')

	return b.String(), true
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	s, _ := ScanString(br, '\n')
	ans, ok := solve(s)
	if !ok {
		bw.WriteString("Impossible")
	} else {
		bw.WriteString(ans)
	}
	bw.WriteByte('\n')
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
