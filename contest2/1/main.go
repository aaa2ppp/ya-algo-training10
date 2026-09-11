package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(k int, a []int) int

func solve(k int, a []int) int {
	freq := make(map[int]int)
	for _, num := range a {
		freq[num]++
	}

	var count int

	m := make(map[int]bool)
	for num := range freq {
		if num*2 == k && freq[num] > 1 {
			count += freq[num] - 1
			continue
		}
		if m[k-num] {
			count += min(freq[num], freq[k-num])
			continue
		}
		m[num] = true
	}

	return count
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, k int
	ScanInt(br, &n, &k)

	a := make([]int, n)
	ScanInts(br, a)

	ans := solve(k, a)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
