package main

import "fmt"

type Point struct {
	X, Y int
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func PointsP1() ([]Point, int) {
	rows := ReadInputRows()
	resp := make([]Point, len(rows)+1)
	boundary := 0

	d := map[byte][2]int{
		'R': {1, 0},
		'D': {0, 1},
		'L': {-1, 0},
		'U': {0, -1},
	}

	p := Point{
		X: 0,
		Y: 0,
	}
	resp[0] = p

	var val int
	for i, row := range rows {
		if row[3] == ' ' {
			val = Atoi(row[2:3])
		} else {
			val = Atoi(row[2:4])
		}

		boundary += val
		diff := d[row[0]]

		p = Point{
			X: p.X + (val * diff[0]),
			Y: p.Y + (val * diff[1]),
		}

		resp[i+1] = p
	}

	return resp, boundary
}

func PointsP2() ([]Point, int) {
	rows := ReadInputRows()
	resp := make([]Point, len(rows)+1)
	boundary := 0

	d := map[byte][2]int{
		'0': {1, 0},
		'1': {0, 1},
		'2': {-1, 0},
		'3': {0, -1},
	}

	p := Point{
		X: 0,
		Y: 0,
	}
	resp[0] = p

	for i, row := range rows {
		val := HexToInt(row[len(row)-7 : len(row)-2])
		boundary += val
		diff := d[row[len(row)-2]]
		p = Point{
			X: p.X + (val * diff[0]),
			Y: p.Y + (val * diff[1]),
		}

		resp[i+1] = p
	}

	return resp, boundary
}

func CalulateArea(points []Point, boundary int) int {
	var A int = 0

	// https://en.wikipedia.org/wiki/Shoelace_formula
	for i, point := range points {
		A += point.X * (points[mod(i-1, len(points))].Y - points[mod(i+1, len(points))].Y)
	}

	if A < 0 {
		A = -A
	}

	A /= 2

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	return (A - boundary/2 + 1) + boundary
}
func main() {
	p1P, p1B := PointsP1()
	p2P, p2B := PointsP2()
	fmt.Println("Puzzle One:", CalulateArea(p1P, p1B))
	fmt.Println("Puzzle Two:", CalulateArea(p2P, p2B))
}
