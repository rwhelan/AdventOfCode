package main

import (
	"fmt"
	"slices"
)

func FirstDigit(b []byte) int {
	for _, c := range b {
		if c >= 48 && c <= 57 {
			return int(c) - 48
		}
	}

	panic("No digit found")
}

func LastDigit(input []byte) int {
	b := make([]byte, len(input))
	copy(b, input)

	slices.Reverse(b)
	return FirstDigit(b)
}

func PuzzleOne() int {
	total := 0
	for _, row := range ReadInputRows() {
		total += (FirstDigit(row) * 10) + LastDigit(row)
	}

	return total
}

func main() {
	fmt.Printf("Puzzle One: %d\n", PuzzleOne())
}
