package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Race struct {
	time int
	dist int
}

func (r Race) WaysToWin() int {
	t := 0
	for i := 1; i < r.time; i++ {
		if (r.time-i)*i > r.dist {
			t++
		}
	}

	return t
}

func Parse() []Race {
	file := ReadInputRows()
	var times []int
	var dists []int

	for i, c := range file[0] {
		if IsDigit(c) {
			times = SplitInts(file[0][i:])
			break
		}
	}

	for i, c := range file[1] {
		if IsDigit(c) {
			dists = SplitInts(file[1][i:])
			break
		}
	}

	if len(times) != len(dists) {
		panic("Bad file format")
	}

	resp := make([]Race, len(times))
	for i := 0; i < len(times); i++ {
		resp[i] = Race{
			time: times[i],
			dist: dists[i],
		}
	}

	return resp
}

func ConcatInts(i []int) int {
	s := bytes.Buffer{}
	for _, v := range i {
		s.WriteString(strconv.Itoa(v))
	}

	return Atoi(s.Bytes())
}

func main() {
	races := Parse()
	t := 1

	for _, race := range races {
		t *= race.WaysToWin()
	}

	fmt.Println("Puzzle One: ", t)

	times := make([]int, len(races))
	dist := make([]int, len(races))

	for i := 0; i < len(races); i++ {
		times[i] = races[i].time
		dist[i] = races[i].dist
	}

	race := Race{
		time: ConcatInts(times),
		dist: ConcatInts(dist),
	}

	fmt.Println("Puzzle Two: ", race.WaysToWin())
}
