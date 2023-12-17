package main

import (
	"bytes"
	"fmt"
	"slices"
)

var Numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func FirstDigit(b []byte) (int, int) {
	for i, c := range b {
		if c >= 48 && c <= 57 {
			return int(c) - 48, i
		}
	}

	panic("No digit found")
}

func LastDigit(input []byte) (int, int) {
	b := make([]byte, len(input))
	copy(b, input)

	slices.Reverse(b)
	return FirstDigit(b)
}

func FirstNumber(b []byte) (int, int) {
	i := 255
	n := -1

	for k, v := range Numbers {
		idx := bytes.Index(b, []byte(k))
		if idx < i && idx != -1 {
			i = idx
			n = v
		}
	}

	return n, i
}

func LastNumber(input []byte) (int, int) {
	i := 255
	n := -1

	b := make([]byte, len(input))
	copy(b, input)
	slices.Reverse(b)

	for s, v := range Numbers {
		k := []byte(s)
		slices.Reverse(k)
		idx := bytes.Index(b, k)
		if idx < i && idx != -1 {
			i = idx
			n = v
		}
	}

	return n, i
}

func PuzzleOne() int {
	total := 0
	for _, row := range ReadInputRows() {
		f, _ := FirstDigit(row)
		l, _ := LastDigit(row)
		total += (f * 10) + l
	}

	return total
}

func PuzzleTwo() int {
	total := 0
	for _, row := range ReadInputRows() {
		fd, fdI := FirstDigit(row)
		ld, ldI := LastDigit(row)
		fn, fnI := FirstNumber(row)
		ln, lnI := LastNumber(row)

		f, l := fd, ld
		if fnI < fdI {
			f = fn
		}

		if lnI < ldI {
			l = ln
		}

		total += (f * 10) + l
	}

	return total
}

func main() {
	fmt.Printf("Puzzle One: %d\n", PuzzleOne())
	fmt.Printf("Puzzle Two: %d\n", PuzzleTwo())
}
