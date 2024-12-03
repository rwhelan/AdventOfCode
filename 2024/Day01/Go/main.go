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

	total := 0
	for i := range colOne {
		total += int(math.Abs(float64(colOne[i] - colTwo[i])))
	}

	fmt.Println(total)
}
