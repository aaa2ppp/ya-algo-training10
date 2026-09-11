package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/aaa2ppp/contestio"
)

func Test_run_solve(t *testing.T) {
	test_run(t, solve)
}

func test_run(t *testing.T, solve solveFunc) {
	tests := []struct {
		name    string
		input   string
		wantOut string
		debug   bool
	}{
		{
			"1",
			`<test></test>
`,
			`<test></test>
`,
			true,
		},
		{
			"2",
			`test<tist>/<>
`,
			`Impossible`,
			true,
		},
		{
			"3",
			`te<ste>st/<t>
`,
			`<test></test>
`,
			true,
		},
		{
			"4",
			`<>test<>//<>test<>
`,
			`<e><stt></stt></e>
`,
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debug = v }(debug)
			debug = tt.debug

			in := strings.NewReader(tt.input)
			out := &bytes.Buffer{}
			run(in, out, solve)
			gotOut := out.String()
			if len(gotOut) > 0 && gotOut[0] == '<' {
				if !checkXML(gotOut) {
					t.Errorf("run() = %v, invalid XML", gotOut)
				}
				if !reflect.DeepEqual(countChars(gotOut), countChars(tt.wantOut)) {
					t.Errorf("run() = %v, want char set %v", gotOut, tt.wantOut)
				}
			} else if trimLines(gotOut) != trimLines(tt.wantOut) {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func trimLines(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t\r\n")
	}
	for n := len(lines); n > 0 && lines[n-1] == ""; n-- {
		lines = lines[:n-1]
	}
	return strings.Join(lines, "\n")
}

func countChars(s string) map[byte]int {
	s = strings.TrimSpace(s)
	m := make(map[byte]int, 26+3)
	for _, c := range []byte(s) {
		m[c]++
	}
	return m
}

func checkXML(s string) bool {
	s = strings.TrimSpace(s)
	var st contestio.Stack[string]

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

			if st.Pop() != s[begin:end] {
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
