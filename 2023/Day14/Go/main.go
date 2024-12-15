package main

import (
	"fmt"
	"slices"
	"strings"
)

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
	for i, b := range column {
		if b == 'O' {
			total += len(column) - i
		}
	}

	return total
}

func TiltGridRight(g Array2D) Array2D {
	for _, row := range g {
		for i := 0; i < len(row); {
			for j := i; ; j++ {
				if j == len(row) || row[j] == '#' {
					slices.Sort(row[i:j])
					i = j + 1
					break
				}
			}
		}
	}

	return g
}

func CycleGrid(g Array2D) Array2D {
	for i := 0; i < 4; i++ {
		g = g.RotateCW()
		g = TiltGridRight(g)
	}

	return g
}

func GridLoad(g Array2D) int {
	load := 0
	for i := range g[0] {
		load += ColumnValue(ReadColumn(g, i))
	}

	return load
}

func main() {
	grid := Parse()

	p1g := grid.RotateCW()
	p1g = TiltGridRight(p1g)
	p1g = p1g.RotateCCW()

	fmt.Println("Puzzle One:", GridLoad(p1g))

	var cur int
	var start int
	gridList := make([]string, 0, 100000)
	seenGrids := make(map[string]Array2D)

	seenGrids[grid.Hash()] = grid.Clone()
	gridList = append(gridList, grid.Hash())

	for i := 1; ; i++ {
		grid = CycleGrid(grid)
		hash := grid.Hash()
		if _, ok := seenGrids[hash]; ok {
			cur = i
			break
		}
		seenGrids[hash] = grid.Clone()
		gridList = append(gridList, hash)

	}

	for i, gh := range gridList {
		if gh == grid.Hash() {
			start = i
			break
		}
	}

	v := (1000000000-start)%(cur-start) + start
	fmt.Println("Puzzle Two:", GridLoad(seenGrids[gridList[v]]))
}

// ==============================================================

func P1ColumnValue(column []byte) int {
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

func Pgrid(b Array2D) {
	for _, row := range b {
		Prow(row)
	}
}

func Prow(b []byte) {
	s := strings.Builder{}
	for _, c := range b {
		s.WriteByte(c)
	}

	fmt.Println(s.String())
}
