package main

import "fmt"

type coord []int
type coordlist []coord

func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func calcDist(l, r coord) int {
	return Abs(l[0]-r[0]) + Abs(l[1]-r[1])
}

func AddDists(M Map2D) int {
	clist := make(coordlist, 0)

	for y, row := range M {
		for x, b := range row {
			if b == '#' {
				clist = append(clist, coord{y, x})
			}
		}
	}

	distance := 0
	for i := 0; i < len(clist); i++ {
		for j := i + 1; j < len(clist); j++ {
			distance += calcDist(clist[i], clist[j])
		}
	}

	return distance
}

func main() {
	M := ReadInputMap()

	Mp1 := Expand(M, 2)
	Mp2 := Expand(M, 1000000)

	fmt.Println("Puzzle One:", AddDists(Mp1))
	fmt.Println("Puzzle Two:", AddDists(Mp2))
}
