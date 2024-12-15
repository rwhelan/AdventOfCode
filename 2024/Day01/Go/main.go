package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	rawRows := ReadInputRows()

	colOne := make([]int, len(rawRows))
	colTwo := make([]int, len(rawRows))

	for i := range rawRows {
		colOne[i], _ = strconv.Atoi(string(rawRows[i][0:5]))
		colTwo[i], _ = strconv.Atoi(string(rawRows[i][8:13]))
	}

	sort.Ints(colOne)
	sort.Ints(colTwo)

	pOneTotal := 0
	pTwoLookup := make(map[int]int)
	for i, num := range colTwo {
		pOneTotal += int(math.Abs(float64(colOne[i] - colTwo[i])))
		pTwoLookup[num]++
	}

	fmt.Println("Problem One:", pOneTotal)

	pTwoTotal := 0
	for _, num := range colOne {
		pTwoTotal += num * pTwoLookup[num]
	}

	fmt.Println("Problem Two:", pTwoTotal)

}
