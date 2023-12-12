package main

import "fmt"

func Expand(M Map2D) Map2D {
	resp := make(Map2D)
	cols := make(map[int]bool)

	offSet := 0
	for y := 0; y < len(M); y++ {
		row := M[y]
		resp[y+offSet] = make(map[int]byte)
		dupRow := true
		for x := 0; x < len(row); x++ {
			if _, ok := cols[x]; !ok {
				cols[x] = true
			}

			b := M[y][x]
			resp[y+offSet][x] = b
			if b != '.' {
				cols[x] = false
				dupRow = false
			}
		}

		if dupRow {
			offSet++
			fmt.Println(y, y+offSet)
			resp[y+offSet] = make(map[int]byte)
			for i := 0; i < len(row); i++ {
				resp[y+offSet][i] = '.'
			}
		}
	}

	return resp
}

func PrintMap(M Map2D) {
	for y := 0; y < len(M); y++ {
		row := M[y]
		for x := 0; x < len(row); x++ {
			fmt.Print(string(M[y][x]))
		}
		fmt.Println()
	}
}

func main() {
	M := ReadInputMap()
	eM := Expand(M)

	PrintMap(M)
	fmt.Println()
	PrintMap(eM)
}
