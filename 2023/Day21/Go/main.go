package main

import (
	"fmt"
)

type Node struct {
	Y, X    int
	Visited bool
	Rock    bool
	v       byte
}

type Grid [][]*Node

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

func (g Grid) Count(b byte) int {
	total := 0

	for _, row := range g {
		for _, node := range row {
			if node.v == b {
				total++
			}
		}
	}

	return total
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

		resp = append(resp, g[node.Y+yd][node.X+xd])
	}

	return resp
}

type QueueSet struct {
	set   map[[2]int]bool
	queue chan *Node
}

func NewQueueSet() *QueueSet {
	return &QueueSet{
		set:   make(map[[2]int]bool),
		queue: make(chan *Node, 1000000),
	}
}

func (q *QueueSet) AddNode(node *Node) {
	if !q.set[[2]int{node.X, node.Y}] {
		q.queue <- node
		q.set[[2]int{node.X, node.Y}] = true
	}
}

func (q *QueueSet) GetNext() *Node {
	select {
	case node := <-q.queue:
		return node
	default:
		return nil
	}
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
	grid, startNode := ParseMap()
	queue := NewQueueSet()
	queue.AddNode(startNode)

	for i := 0; i <= 64; i++ {
		nS := make([]*Node, 0, 1000)
		for n := queue.GetNext(); n != nil; n = queue.GetNext() {
			n.Visited = true
			n.v = '0'
			nS = append(nS, grid.FindUVNs(n)...)
		}

		for _, node := range nS {
			queue.AddNode(node)
		}

		grid.Flop()
	}

	grid.Flop()
	fmt.Println("Puzzle One:", grid.Count('0'))
}
