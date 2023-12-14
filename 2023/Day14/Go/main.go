package main

import "fmt"

func Parse() Array2D {
	return ReadInputRows()
}

func ReadColumn(g Array2D, c int) []byte {
	resp := make([]byte, len(g))
	for i, row := range g {
		resp[i] = row[c]
	}

	return resp
}

func ColumnValue(c []byte) int {
	ptr, cur, val := 0, 0, 0

	ptrSeek := func() {
		for ; c[ptr] != '.'; ptr++ {
		}
	}

	curSeek := func() bool {
		for ; c[cur] != '0'; cur++ {
			fmt.Println(cur, len(c)-1)
			if cur == len(c)-1 {
				return true
			}
		}

		return false
	}

	for !curSeek() {
		fmt.Println("H")
		ptrSeek()
		val += len(c) - ptr
	}

	return val
}

func main() {
	grid := Parse()

	c := ReadColumn(grid, 1)
	fmt.Println(string(c))
	fmt.Println(ColumnValue(c))
}
