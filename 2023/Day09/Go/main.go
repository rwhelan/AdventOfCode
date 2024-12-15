package main

import (
	"fmt"
	"slices"
)

func IsZeros(numbers []int) bool {
	for _, i := range numbers {
		if i != 0 {
			return false
		}
	}

	return true
}

func Calc(numbers []int) int {
	r := make([][]int, 0)
	Diff(numbers, &r)

	n := r[0][len(r[0])-1]
	for i := 1; i < len(r); i++ {
		n += r[i][len(r[i])-1]
	}

	return numbers[len(numbers)-1] + n
}

func Diff(numbers []int, res *[][]int) {
	if IsZeros(numbers) {
		return
	}

	resp := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		n := numbers[i+1] - numbers[i]
		resp[i] = n
	}

	*res = append(*res, resp)
	Diff(resp, res)
}

func main() {
	p1, p2 := 0, 0
	for _, row := range ReadInputRows() {
		numbers := SplitInts(row)
		p1 += Calc(numbers)
		slices.Reverse(numbers)
		p2 += Calc(numbers)
	}

	fmt.Println("Puzzle One:", p1)
	fmt.Println("Puzzle Two:", p2)
}
