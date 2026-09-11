package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(n int, w, s string) (string, error)

func completed(w string) bool {
	var set byte
	for _, c := range []byte(w) {
		switch c {
		case '(':
			set |= 1
		case ')':
			set |= 2
		case '[':
			set |= 4
		case ']':
			set |= 8
		}
	}
	return set == 15
}

func solve(n int, w, s string) (string, error) {
	if n%2 != 0 || len(s) > n {
		return "", fmt.Errorf("no solution for n=%d s=%q", n, s)
	}
	if !completed(w) {
		return "", fmt.Errorf("incompleted set %q", w)
	}

	var st Stack[byte]
	var ans strings.Builder
	ans.Grow(n)

	for i, c := range []byte(s) {
		switch c {
		case '(', '[':
			st.Push(c)
		case ')':
			if st.Empty() || st.Pop() != '(' {
				return "", fmt.Errorf("%d: unexpected char '%c' (%d)", i+1, c, c)
			}
		case ']':
			if st.Empty() || st.Pop() != '[' {
				return "", fmt.Errorf("%d: unexpected char '%c' (%d)", i+1, c, c)
			}
		default:
			return "", fmt.Errorf("%d: unknown char '%c' (%d)", i+1, c, c)
		}
		n--
		ans.WriteByte(c)
	}

	for n > 0 {
		possibleOpen := n-st.Len() > 1
		var possibleClose byte
		if !st.Empty() {
			switch st.Top() {
			case '(':
				possibleClose = ')'
			case '[':
				possibleClose = ']'
			}
		}
		if !possibleOpen && possibleClose == 0 {
			return "", fmt.Errorf("no solution for n=%d s=%q", n, s)
		}
		for _, c := range []byte(w) {
			if possibleOpen && (c == '(' || c == '[') {
				st.Push(c)
				ans.WriteByte(c)
				break
			}
			if c == possibleClose {
				st.Pop()
				ans.WriteByte(c)
				break
			}
		}
		n--
	}

	return ans.String(), nil
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanIntLn(br, &n)

	w, _ := ScanString(br, '\n')
	s, _ := ScanString(br, '\n')

	ans, err := solve(n, w, s)
	if err != nil {
		panic(err)
	}
	bw.WriteString(ans)
	bw.WriteByte('\n')
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
