package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func IsDigit(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	}

	return false
}

func IsSymbol(b byte) bool {
	return b != '.' && !IsDigit(b)
}

func SymbolIdx(schematic []byte) map[int]map[int]byte {
	resp := make(map[int]map[int]byte)
	for y, row := range bytes.Split(schematic, []byte{'\n'}) {
		resp[y] = make(map[int]byte)
		for x, char := range row {
			if IsSymbol(char) {
				resp[y][x] = char
			}
		}
	}

	return resp
}

func AdjacentSymbol(x, y int, symbols map[int]map[int]byte) bool {
	for ly := -1; ly < 2; ly++ {
		for lx := -1; lx < 2; lx++ {
			if _, ok := symbols[ly+y][lx+x]; ok {
				return true
			}
		}
	}

	return false
}

func PuzzleOne(schematic []byte) int {
	symbols := SymbolIdx(schematic)

	total := 0
	for y, row := range bytes.Split(schematic, []byte{'\n'}) {
		buff := make([]byte, 0, 3)
		startidx, endidx := 0, 0

		for x, c := range row {
			if IsDigit(c) {
				if startidx == 0 {
					startidx = x
				}

				endidx = x
				buff = append(buff, c)

			} else if len(buff) != 0 {
				if AdjacentSymbol(startidx, y, symbols) || AdjacentSymbol(endidx, y, symbols) {
					num, err := strconv.Atoi(string(buff))
					if err != nil {
						panic(err)
					}

					total += num
				}
				startidx = 0
				buff = buff[:0]

			}
		}

		if len(buff) != 0 && AdjacentSymbol(startidx, y, symbols) || AdjacentSymbol(len(row)-1, y, symbols) {
			num, err := strconv.Atoi(string(buff))
			if err != nil {
				panic(err)
			}

			total += num
		}
	}

	return total
}

func IsGear(x, y, l int, symbols map[int]map[int]byte) string {
	for ly := -1; ly < 2; ly += 2 {
		for lx := -(l + 1); lx <= 0; lx++ {
			if c, ok := symbols[ly+y][lx+x]; ok && c == '*' {
				return fmt.Sprintf("%d:%d", ly+y, lx+x)
			}
		}
	}

	if c, ok := symbols[y][x]; ok && c == '*' {
		return fmt.Sprintf("%d:%d", y, x)
	}

	if c, ok := symbols[y][x-(l+1)]; ok && c == '*' {
		return fmt.Sprintf("%d:%d", y, x-(l+1))
	}

	return ""
}

func Atoi(b []byte) int {
	num, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}

	return num
}

func PuzzleTwo(schematic []byte) int {
	symbols := SymbolIdx(schematic)
	gears := make(map[string][]int)

	for y, row := range bytes.Split(schematic, []byte{'\n'}) {
		buff := make([]byte, 0, 3)
		for x, b := range row {
			if IsDigit(b) {
				buff = append(buff, b)
			}

			if len(buff) != 0 && !IsDigit(b) || x == len(row)-1 {
				if g := IsGear(x, y, len(buff), symbols); g != "" {
					_, ok := gears[g]
					if !ok {
						gears[g] = make([]int, 0, 2)
					}

					gears[g] = append(gears[g], Atoi(buff))
				}

				buff = buff[:0]
			}
		}
	}

	total := 0
	for _, gr := range gears {
		if len(gr) == 2 {
			total += gr[0] * gr[1]
		}
	}

	return total
}

func main() {
	schematic := ReadInputFile()
	fmt.Println("Puzzle One: ", PuzzleOne(schematic))
	fmt.Println("Puzzle Two: ", PuzzleTwo(schematic))
}
