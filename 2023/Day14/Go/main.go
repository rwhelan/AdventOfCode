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

func ColumnValue(column []byte) int {
	total := 0

	addRocks := func(c, n int) {
		for i := 0; i < n; i++ {
			total += len(column) - (c + i)
		}
	}

	for i := 0; i < len(column); i++ {
		rocks := 0
		for j, b := range column[i:] {
			if b == 'O' {
				rocks++
			}

			if b == '#' || i+j == len(column)-1 {
				addRocks(i, rocks)
				i = j + i
				break
			}
		}
	}

	return total
}

func main() {
	grid := Parse()

	total := 0
	for i := range grid[0] {
		total += ColumnValue(ReadColumn(grid, i))
	}

	fmt.Println(total)
}
