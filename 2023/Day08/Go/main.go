package main

import "fmt"

type Node struct {
	Current, Left, Right string
}

type Nodes []Node

func (n Nodes) Complete(m map[string]Node) bool {
	for _, node := range n {
		if node.Current[len(node.Current)-1] != 'Z' {
			return false
		}
	}

	return true
}

func (n *Node) Step(d byte, m map[string]Node) Node {
	var nn Node

	switch d {
	case 'R':
		nn = m[n.Right]
	case 'L':
		nn = m[n.Left]
	}

	*n = nn
	return nn
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
			Current: string(row[:3]),
			Left:    string(row[7:10]),
			Right:   string(row[12:15]),
		}
	}

	return directions, resp
}

func PuzzleOne(d []byte, m map[string]Node) int {
	count := 0
	current := m["AAA"]

	for {
		cd := d[count%len(d)]
		current.Step(cd, m)
		count++
		if current.Current == "ZZZ" {
			break
		}
	}

	return count
}

func PuzzleTwo(d []byte, m map[string]Node) int {
	count := 0
	current := make(Nodes, 0)

	for n := range m {
		if n[len(n)-1] == 'A' {
			current = append(current, m[n])
		}
	}

	for {
		cd := d[count%len(d)]

		for i, n := range current {
			current[i] = n.Step(cd, m)
		}
		count++

		if current.Complete(m) {
			break
		}
	}

	return count
}

func Walk(start string, ends []string, d []byte, m map[string]Node) int {
	count := 0
	current := m[start]

outer:
	for {
		cd := d[count%len(d)]
		current.Step(cd, m)
		count++

		for _, e := range ends {
			if current.Current == e {
				fmt.Println(e)
				break outer
			}
		}
	}

	return count
}

func main() {
	d, m := Parse()

	// fmt.Println("Puzzle One:", PuzzleOne(d, m))
	// fmt.Println("Puzzle Two:", PuzzleTwo(d, m))
	for _, s := range []string{"NDA", "AAA", "PTA", "PBA", "DVA", "HCA"} {
		fmt.Println(s, Walk(s, []string{"CMZ", "RLZ", "BLZ", "ZZZ", "LXZ", "PMZ"}, d, m))
	}
}
