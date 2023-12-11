package main

import "fmt"

type Previous int

const (
	ABOVE Previous = iota
	RIGHT
	BELOW
	LEFT
)

type Cursor struct {
	Y, X    int
	Current byte
	M       Map2D
	Prev    Previous
}

func (c *Cursor) ReadAbove() (byte, bool) {
	b, ok := c.M[c.Y-1][c.X]
	return b, ok
}

func (c *Cursor) ReadBelow() (byte, bool) {
	b, ok := c.M[c.Y+1][c.X]
	return b, ok
}

func (c *Cursor) ReadRight() (byte, bool) {
	b, ok := c.M[c.Y][c.X+1]
	return b, ok
}

func (c *Cursor) ReadLeft() (byte, bool) {
	b, ok := c.M[c.Y][c.X-1]
	return b, ok
}

func (c *Cursor) MoveUp() {
	c.Y = c.Y - 1
	c.Current = c.M[c.Y][c.X]
	c.Prev = BELOW
}

func (c *Cursor) MoveDown() {
	c.Y = c.Y + 1
	c.Current = c.M[c.Y][c.X]
	c.Prev = ABOVE
}

func (c *Cursor) MoveRight() {
	c.X = c.X + 1
	c.Current = c.M[c.Y][c.X]
	c.Prev = LEFT
}

func (c *Cursor) MoveLeft() {
	c.X = c.X - 1
	c.Current = c.M[c.Y][c.X]
	c.Prev = RIGHT
}

func NewCursor(M Map2D) *Cursor {
	for y, row := range M {
		for x, c := range row {
			if c == 'S' {
				return &Cursor{
					Y:       y,
					X:       x,
					M:       M,
					Current: M[y][x],
				}
			}
		}
	}

	panic("Start Not Found")
}

func main() {
	C := NewCursor(ReadInputMap())
	C.MoveUp()

	count := 1

walk:
	for {
		switch C.Current {
		case 'S':
			break walk

		case 'J':
			if C.Prev == LEFT {
				C.MoveUp()
				break
			}

			if C.Prev == ABOVE {
				C.MoveLeft()
				break
			}

			fmt.Println("Y:X:P", C.Y, C.X, C.Prev)
			panic("Bad Map?! (J)")

		case 'F':
			if C.Prev == RIGHT {
				C.MoveDown()
				break
			}

			if C.Prev == BELOW {
				C.MoveRight()
				break
			}

			fmt.Println("Y:X:P", C.Y, C.X, C.Prev)
			panic("Bad Map?! (F)")

		case '7':
			if C.Prev == LEFT {
				C.MoveDown()
				break
			}

			if C.Prev == BELOW {
				C.MoveLeft()
				break
			}

			fmt.Println("Y:X:P", C.Y, C.X, C.Prev)
			panic("Bad Map?! (7)")

		case 'L':
			if C.Prev == RIGHT {
				C.MoveUp()
				break
			}

			if C.Prev == ABOVE {
				C.MoveRight()
				break
			}

			fmt.Println("Y:X:P", C.Y, C.X, C.Prev)
			panic("Bad Map?! (L)")

		case '|':
			if C.Prev == ABOVE {
				C.MoveDown()
				break
			}

			if C.Prev == BELOW {
				C.MoveUp()
				break
			}

			fmt.Println("Y:X:P", C.Y+1, C.X+1, C.Prev)
			panic("Bad Map?! (|)")

		case '-':
			if C.Prev == LEFT {
				C.MoveRight()
				break
			}

			if C.Prev == RIGHT {
				C.MoveLeft()
				break
			}

			fmt.Println("Y:X:P", C.Y, C.X, C.Prev)
			panic("Bad Map?! (-)")
		}

		count++
	}

	fmt.Println(count)
}
