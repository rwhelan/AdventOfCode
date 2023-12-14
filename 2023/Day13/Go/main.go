package main

import (
	"bytes"
	"fmt"
)

func ParseArrays() []Array2D {
	c := ReadInputFile()
	if c[len(c)-1] == '\n' {
		c = c[:len(c)-1]
	}

	arrays := bytes.Split(c, []byte{'\n', '\n'})

	resp := make([]Array2D, len(arrays))
	for i, a := range arrays {
		resp[i] = Array2D(bytes.Split(a, []byte{'\n'}))
	}

	return resp
}

func RowDiff[T byte | rune | int](left, right []T) int {
	if len(left) != len(right) {
		panic("Arrays not equal Length")
	}

	d := 0
	for i := range left {
		if left[i] != right[i] {
			d++
		}
	}

	return d
}

func FindHorizontalMirror(a Array2D, diff int) int {
outter:
	for r := 0; r < len(a); r++ {
		fdiff := 0
		for o := 0; r-o >= 0 && r+o+1 < len(a); o++ {
			aboveRow := a[r-o]
			belowRow := a[r+o+1]
			fdiff += RowDiff(aboveRow, belowRow)
			if fdiff > diff {
				continue outter
			}
		}

		if fdiff != diff {
			continue
		}

		if r == len(a)-1 {
			return -1
		}

		return r
	}

	return -1
}

func HorizontalScore(a Array2D, diff int) int {
	row := FindHorizontalMirror(a, diff)
	if row != -1 {
		return (1 + row) * 100
	}

	return 0
}

func VerticalScore(a Array2D, diff int) int {
	va := a.RotateCW()
	col := FindHorizontalMirror(va, diff)
	if col != -1 {
		return 1 + col
	}

	return 0
}

func ArrayScore(a Array2D, diff int) int {
	hs := HorizontalScore(a, diff)
	vs := VerticalScore(a, diff)

	if hs > vs {
		return hs
	}

	return vs
}
func main() {
	a := ParseArrays()

	p1total := 0
	p2total := 0

	for i := 0; i < len(a); i++ {
		p1total += ArrayScore(a[i], 0)
		p2total += ArrayScore(a[i], 1)
	}

	fmt.Println("Puzzle One:", p1total)
	fmt.Println("Puzzle Two:", p2total)
}
