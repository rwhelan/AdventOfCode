package main

import (
	"fmt"
	"sort"
)

type DIRECTION int

const (
	UP DIRECTION = iota
	RIGHT
	DOWN
	LEFT
)

type Node struct {
	X, Y int
	D    DIRECTION
}

func CalulateBeam(grid Array2D, X, Y int, D DIRECTION) int {
	seen := make(map[Node]bool)
	count := make(map[[2]int]bool)
	queue := make(chan Node, 100)

	queue <- Node{X: X, Y: Y, D: D}

	for len(queue) != 0 {
		c := <-queue
		for {
			if c.Y > len(grid)-1 || c.X > len(grid[0])-1 || c.Y < 0 || c.X < 0 {
				break
			}

			if seen[c] {
				break
			}

			seen[c] = true
			count[[2]int{c.X, c.Y}] = true

			switch grid[c.Y][c.X] {
			case '\\':
				switch c.D {
				case UP:
					c = Node{X: c.X - 1, Y: c.Y, D: LEFT}
					continue
				case LEFT:
					c = Node{X: c.X, Y: c.Y - 1, D: UP}
					continue
				case DOWN:
					c = Node{X: c.X + 1, Y: c.Y, D: RIGHT}
					continue
				case RIGHT:
					c = Node{X: c.X, Y: c.Y + 1, D: DOWN}
					continue
				}

			case '/':
				switch c.D {
				case UP:
					c = Node{X: c.X + 1, Y: c.Y, D: RIGHT}
					continue
				case LEFT:
					c = Node{X: c.X, Y: c.Y + 1, D: DOWN}
					continue
				case DOWN:
					c = Node{X: c.X - 1, Y: c.Y, D: LEFT}
					continue
				case RIGHT:
					c = Node{X: c.X, Y: c.Y - 1, D: UP}
					continue
				}

			case '|':
				switch c.D {
				case LEFT, RIGHT:
					queue <- Node{X: c.X, Y: c.Y + 1, D: DOWN}
					c = Node{X: c.X, Y: c.Y - 1, D: UP}
					continue
				}

			case '-':
				switch c.D {
				case UP, DOWN:
					queue <- Node{X: c.X + 1, Y: c.Y, D: RIGHT}
					c = Node{X: c.X - 1, Y: c.Y, D: LEFT}
					continue
				}
			}

			switch c.D {
			case RIGHT:
				c = Node{X: c.X + 1, Y: c.Y, D: RIGHT}
			case LEFT:
				c = Node{X: c.X - 1, Y: c.Y, D: LEFT}
			case UP:
				c = Node{X: c.X, Y: c.Y - 1, D: UP}
			case DOWN:
				c = Node{X: c.X, Y: c.Y + 1, D: DOWN}
			}
		}
	}

	return len(count)
}
func main() {
	grid := Array2D(ReadInputRows())

	fmt.Println("Puzzle One:", CalulateBeam(grid, 0, 0, RIGHT))

	pathLens := make([]int, 0, len(grid)*2+len(grid[0])*2)
	for i := 0; i < len(grid[0]); i++ {
		pathLens = append(pathLens, CalulateBeam(grid, i, 0, DOWN))
		pathLens = append(pathLens, CalulateBeam(grid, i, len(grid)-1, UP))
	}
	for i := 0; i < len(grid); i++ {
		pathLens = append(pathLens, CalulateBeam(grid, 0, i, RIGHT))
		pathLens = append(pathLens, CalulateBeam(grid, len(grid[0])-1, i, LEFT))
	}

	sort.Ints(pathLens)
	fmt.Println("Puzzle Two:", pathLens[len(pathLens)-1])
}
