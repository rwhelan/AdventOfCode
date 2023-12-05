package main

import (
	"bytes"
	"fmt"
	"slices"
)

type Range struct {
	dst, src, rng int
}

func (r Range) Find(i int) (int, bool) {
	if i >= r.src && i < (r.src+r.rng) {
		return r.dst + (i - r.src), true
	}

	return -1, false
}

type Map struct {
	dst, src string
	ranges   []Range
}

func (m Map) Lookup(i int) int {
	for _, r := range m.ranges {
		if resp, ok := r.Find(i); ok {
			return resp
		}
	}

	return i
}

func Parse() ([]int, map[string]Map) {
	var seeds []int
	var currentMap string
	maps := make(map[string]Map)

	names := func(row []byte) (string, string) {
		pieces := bytes.Split(row[:len(row)-5], []byte{'-'})
		return string(pieces[0]), string(pieces[2])
	}

	for i, row := range ReadInputRows() {
		if len(row) == 0 {
			continue

		} else if row[len(row)-1] == ':' {
			src, dst := names(row)
			maps[src] = Map{
				src:    src,
				dst:    dst,
				ranges: make([]Range, 0, 1000),
			}

			currentMap = src

		} else if IsDigit(row[0]) {
			vals := SplitInts(row)
			if len(vals) != 3 {
				fmt.Println("ROW :", i)
				panic("Bad File Format")
			}

			cm := maps[currentMap]
			cm.ranges = append(cm.ranges, Range{
				vals[0],
				vals[1],
				vals[2],
			})

			maps[currentMap] = cm

		} else if string(row[:7]) == "seeds: " {
			seeds = SplitInts(row[7:])

		}
	}

	return seeds, maps
}

func Walk(s int, m map[string]Map) int {
	currentMap, ok := m["seed"]
	for ; ok; currentMap, ok = m[currentMap.dst] {
		s = currentMap.Lookup(s)
	}

	return s
}

func main() {
	seeds, m := Parse()
	locations := make([]int, len(seeds))

	for i, s := range seeds {
		locations[i] = Walk(s, m)
	}

	slices.Sort(locations)

	fmt.Println("Puzzle One:", locations[0])
}
