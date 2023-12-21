package main

import "fmt"

type Node struct {
	Y, X    int
	Visited bool
	Rock    bool
	v       byte
}

type Grid [][]*Node

var seen map[[2]int]bool

func (g Grid) Print() {
	for _, row := range g {
		for _, node := range row {
			fmt.Print(string(node.v))
		}

		fmt.Println()
	}
}

func (g Grid) Flop() {
	for _, row := range g {
		for _, node := range row {
			if node.Visited {
				if node.v == '0' {
					node.v = '.'
					continue
				}

				if node.v == '.' {
					node.v = '0'
				}
			}
		}
	}
}

func (g Grid) FindUVNs(node *Node) []*Node {
	resp := make([]*Node, 0, 4)
	for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		yd, xd := d[0], d[1]
		if node.Y+yd < 0 || node.Y+yd > len(g)-1 {
			continue
		}

		if node.X+xd < 0 || node.X+xd > len(g[0])-1 {
			continue
		}

		if g[node.Y+yd][node.X+xd].Visited || g[node.Y+yd][node.X+xd].Rock {
			continue
		}

		if _, ok := seen[[2]int{node.Y + yd, node.X + xd}]; !ok {
			resp = append(resp, g[node.Y+yd][node.X+xd])
			seen[[2]int{node.Y + yd, node.X + xd}] = true
		}
	}

	return resp
}

func ParseMap() (Grid, *Node) {
	rows := ReadInputRows()
	var startNode *Node
	resp := make([][]*Node, len(rows))
	for y, row := range ReadInputRows() {
		resp[y] = make([]*Node, len(rows[0]))

		for x, b := range row {
			resp[y][x] = &Node{
				Y:       y,
				X:       x,
				Visited: false,
				Rock:    b == '#',
				v:       b,
			}

			if resp[y][x].v == 'S' {
				startNode = resp[y][x]
			}
		}
	}

	return resp, startNode
}

func main() {
	seen = make(map[[2]int]bool)

	grid, startNode := ParseMap()
	q := grid.FindUVNs(startNode)

	for i := 0; i < 64000; i++ {
		nq := make([]*Node, 0)
		for _, n := range q {
			n.Visited = true
			n.v = '0'
			nq = append(nq, grid.FindUVNs(n)...)
		}

		grid.Flop()

		q = nq
	}

	grid.Flop()

	total := 0
	for s := range seen {
		if grid[s[0]][s[1]].v == '0' {
			total++
		}
	}

	grid.Print()
	fmt.Println(total)
}
