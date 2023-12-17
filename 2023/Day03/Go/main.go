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

func SymbolIdx(schematic []byte) map[int]map[int]struct{} {
	resp := make(map[int]map[int]struct{})
	for y, row := range bytes.Split(schematic, []byte{'\n'}) {
		resp[y] = make(map[int]struct{})
		for x, char := range row {
			if IsSymbol(char) {
				resp[y][x] = struct{}{}
			}
		}
	}

	return resp
}

func AdjacentSymbol(x, y int, symbols map[int]map[int]struct{}) bool {
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

func main() {
	schematic := ReadInputFile()
	fmt.Println("Puzzle One: ", PuzzleOne(schematic))
}
