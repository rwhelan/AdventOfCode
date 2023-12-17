package main

import (
	"fmt"
	"strconv"
)

type Set struct {
	Red   int
	Green int
	Blue  int
}

func Atoi(b []byte) int {
	num, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}

	return num
}

func GameSets(row []byte) (int, []Set) {
	idx := 5
	buff := make([]byte, 0, 3)
	sets := make([]Set, 0, 10)

	for ; row[idx] != ':'; idx++ {
		buff = append(buff, row[idx])
	}
	idx += 2

	gameNumber := Atoi(buff)
	buff = buff[:0]

	for idx < len(row) {
		s := Set{}
		for {
			for ; row[idx] != ' '; idx++ {
				buff = append(buff, row[idx])
			}
			idx++

			num := Atoi(buff)
			buff = buff[:0]

			switch row[idx] {
			case 'g':
				s.Green = num
				idx += 5

			case 'r':
				s.Red = num
				idx += 3

			case 'b':
				s.Blue = num
				idx += 4

			}

			if idx > len(row)-1 {
				sets = append(sets, s)
				break
			}

			if row[idx] == ';' {
				sets = append(sets, s)
				s = Set{}
			}

			idx += 2
		}
	}

	return gameNumber, sets
}

func PowerSet(sets []Set) int {
	red, blue, green := 0, 0, 0
	for _, set := range sets {
		if set.Red > red {
			red = set.Red
		}

		if set.Green > green {
			green = set.Green
		}

		if set.Blue > blue {
			blue = set.Blue
		}
	}

	if red == 0 || green == 0 || blue == 0 {
		panic("zero")
	}

	return red * green * blue
}

func GamePossible(sets []Set) bool {
	for _, s := range sets {
		if s.Red > 12 || s.Green > 13 || s.Blue > 14 {
			return false
		}
	}

	return true
}

func main() {
	puzzleOne := 0
	puzzleTwo := 0

	for _, row := range ReadInputRows() {
		gameNumber, sets := GameSets(row)
		if GamePossible(sets) {
			puzzleOne += gameNumber
		}

		puzzleTwo += PowerSet(sets)
	}

	fmt.Println("Puzzle One: ", puzzleOne)
	fmt.Println("Puzzle Two: ", puzzleTwo)
}
