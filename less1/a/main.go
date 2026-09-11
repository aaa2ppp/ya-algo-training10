package main

import (
	"fmt"
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string) bool

func solve(s string) bool {
	var st Stack[byte]
	for _, c := range []byte(s) {
		switch c {
		case '(', '[', '{':
			st.Push(c)
		case ')':
			if st.Empty() || st.Pop() != '(' {
				return false
			}
		case ']':
			if st.Empty() || st.Pop() != '[' {
				return false
			}
		case '}':
			if st.Empty() || st.Pop() != '{' {
				return false
			}
		default:
			panic(fmt.Errorf("unknown char '%c' (%d)", c, c))
		}
	}
	return st.Len() == 0
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	s, _ := ScanString(br, '\n')
	if solve(s) {
		bw.WriteString("yes\n")
	} else {
		bw.WriteString("no\n")
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
