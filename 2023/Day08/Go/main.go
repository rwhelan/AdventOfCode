package main

import "fmt"

type Node struct {
	Left, Right string
}

func Parse() ([]byte, map[string]Node) {
	var directions []byte
	resp := make(map[string]Node)
	for i, row := range ReadInputRows() {
		if i == 0 {
			directions = row[:]
			continue
		}

		if i == 1 {
			continue
		}

		resp[string(row[:3])] = Node{
			Left:  string(row[7:10]),
			Right: string(row[12:15]),
		}
	}

	return directions, resp
}

func main() {
	d, m := Parse()
	count := 0
	current := m["AAA"]
	var next string

	for {
		cd := d[count%len(d)]
		if cd == 'L' {
			next = current.Left
		} else {
			next = current.Right
		}

		if next == "ZZZ" {
			break
		}
		current = m[next]
		count++
	}

	fmt.Println("Puzzle One:", count+1)
}
