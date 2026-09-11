package main

import (
	"bytes"
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]byte) []byte

func checkXML(s []byte) bool {
	var st Stack[[]byte]

	end := 0
	for end < len(s) {
		if s[end] != '<' {
			return false
		}
		if end++; end == len(s) {
			return false
		}

		if s[end] == '/' {
			if st.Empty() {
				return false
			}

			if end++; end == len(s) {
				return false
			}

			begin := end
			for end < len(s) && s[end] != '>' {
				if c := s[end]; !('a' <= c && c <= 'z') {
					return false
				}
				end++
			}
			if end == len(s) {
				return false
			}

			if !bytes.Equal(st.Pop(), s[begin:end]) {
				return false
			}
		} else {
			begin := end
			for end < len(s) && s[end] != '>' {
				if c := s[end]; !('a' <= c && c <= 'z') {
					return false
				}
				end++
			}
			if end == len(s) {
				return false
			}

			st.Push(s[begin:end])
		}

		end++
	}

	return st.Len() == 0
}

func slowSolve(s []byte) []byte {
	ab := make([]byte, 0, 26+3)

	ab = append(ab, "</>"...)

	for i := 0; i < 26; i++ {
		ab = append(ab, byte('a'+i))
	}

	for i := range s {
		back := s[i]
		for _, c := range ab {
			s[i] = c
			if checkXML(s) {
				return s
			}
		}
		s[i] = back
	}

	return nil
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	s, _ := ScanBytes(br, '\n')
	ans := solve(s)

	bw.Write(ans)
	bw.WriteByte('\n')
}

func main() {
	run(os.Stdin, os.Stdout, slowSolve)
}
