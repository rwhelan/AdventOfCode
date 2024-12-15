package main

import "fmt"

func Expand(M Map2D, size int) Map2D {
	vMap := make(Map2D)
	cols := make(map[int]bool)

	maxH, maxW := MapHW(M)
	vOffSet := 0
	for y := 0; y < maxH; y++ {
		row := M[y]
		vMap[y+vOffSet] = make(map[int]byte)
		dupRow := true
		for x := 0; x < maxW; x++ {
			if _, ok := cols[x]; !ok {
				cols[x] = true
			}

			b := row[x]
			if b == '#' {
				vMap[y+vOffSet][x] = b
				cols[x] = false
				dupRow = false
			}
		}

		if dupRow {
			vOffSet += size - 1
		}
	}

	addCols := make(map[int]int)
	colCnt := 0
	for i := 0; i < len(cols); i++ {
		if cols[i] {
			colCnt += size - 1
		}

		addCols[i] = colCnt
	}

	hMap := make(Map2D)
	for y, row := range vMap {
		hMap[y] = make(map[int]byte)
		for x, b := range row {
			hMap[y][x+addCols[x]] = b
		}

	}

	return hMap
}

func MapHW(M Map2D) (int, int) {
	var maxWidth int
	var maxHeight int

	for y, row := range M {
		if y > maxHeight {
			maxHeight = y
		}

		for x := range row {
			if x > maxWidth {
				maxWidth = x
			}
		}
	}

	return maxHeight + 1, maxWidth + 1
}

func PrintMap(M Map2D) {
	maxHeight, maxWidth := MapHW(M)

	for y := 0; y < maxHeight; y++ {
		_, ok := M[y]
		if !ok {
			fmt.Println("|")
			continue
		}

		for x := 0; x < maxWidth; x++ {
			b, ok := M[y][x]
			if !ok {
				fmt.Print("|")
			} else {
				fmt.Print(string(b))
			}
		}
		fmt.Println()
	}
}
