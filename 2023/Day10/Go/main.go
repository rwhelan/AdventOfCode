package main

import "fmt"

type Previous int

const (
	_NONE Previous = iota
	ABOVE
	RIGHT
	BELOW
	LEFT
)

type Path struct {
	P map[int]map[int]struct{}
}

func NewPath() Path {
	return Path{
		P: make(map[int]map[int]struct{}),
	}
}

func (p Path) Add(Y, X int) {
	if _, ok := p.P[Y]; !ok {
		p.P[Y] = make(map[int]struct{})
	}

	p.P[Y][X] = struct{}{}
}

func (p Path) IsPath(Y, X int) bool {
	_, ok := p.P[Y][X]
	return ok
}

type Cursor struct {
	Y, X    int
	Current byte
	M       Map2D
	Prev    Previous
	Path    Path
}

func (c *Cursor) update() {
	c.Current = c.M[c.Y][c.X]
	c.Path.Add(c.Y, c.X)
}

func (c *Cursor) MoveUp() {
	c.Y = c.Y - 1
	c.update()
	c.Prev = BELOW
}

func (c *Cursor) MoveDown() {
	c.Y = c.Y + 1
	c.update()
	c.Prev = ABOVE
}

func (c *Cursor) MoveRight() {
	c.X = c.X + 1
	c.update()
	c.Prev = LEFT
}

func (c *Cursor) MoveLeft() {
	c.X = c.X - 1
	c.update()
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
					Path:    NewPath(),
				}
			}
		}
	}

	panic("Start Not Found")
}

type Scanner struct {
	count     int
	inside    bool
	online    bool
	lineEntry Previous
}

func NewScanner() Scanner {
	return Scanner{
		inside: false,
	}
}

func (s Scanner) CountInside(M Map2D, path Path) int {
	for i := 0; i < len(M); i++ {
		y := i
		row := M[i]
		s.inside = false
		for j := 0; j < len(row); j++ {
			x := j
			b := row[j]
			found := false
			switch {
			case !path.IsPath(y, x) && s.inside:
				found = true
				s.count++

			case path.IsPath(y, x) && b == '|':
				s.inside = !s.inside

			case path.IsPath(y, x) && b == 'F':
				s.online = true
				s.lineEntry = BELOW

			case path.IsPath(y, x) && b == 'L':
				s.online = true
				s.lineEntry = ABOVE

			case path.IsPath(y, x) && s.online == true && s.lineEntry == BELOW && b == 'J':
				s.online = false
				s.inside = !s.inside

			case path.IsPath(y, x) && s.online == true && s.lineEntry == ABOVE && b == 'J':
				s.online = false

			case path.IsPath(y, x) && s.online == true && s.lineEntry == ABOVE && b == '7':
				s.online = false
				s.inside = !s.inside

			case path.IsPath(y, x) && s.online == true && s.lineEntry == BELOW && b == '7':
				s.online = false

			}

			if s.inside {
				fmt.Print("\033[31;1m")
			} else {
				fmt.Print("\033[0m")
			}
			switch {
			case path.IsPath(y, x):
				fmt.Print(string(b))

			case found:
				fmt.Printf("\033[32;1m%s", string(b))

			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	return s.count
}

func main() {
	inputMap := ReadInputMap()
	C := NewCursor(inputMap)
	C.MoveUp()
	// C.MoveDown()

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

	fmt.Println("Puzzle One: ", count/2)

	scanner := NewScanner()
	fmt.Println("Puzzle Two: ", scanner.CountInside(inputMap, C.Path))
}
