package main

import "fmt"

type Part struct {
	X, M, A, S int
}

func ParsePart(b []byte) Part {
	p := Part{}
	for i := 0; i < len(b); i++ {
		if b[i] == '=' {
			c := b[i-1]
			i++
			j := i
			for ; ; j++ {
				if !IsDigit(b[j]) {
					break
				}
			}

			switch c {
			case 'x':
				p.X = Atoi(b[i:j])

			case 'm':
				p.M = Atoi(b[i:j])

			case 'a':
				p.A = Atoi(b[i:j])

			case 's':
				p.S = Atoi(b[i:j])

			default:
				panic("Unsure what to do here")

			}

			i = j
		}
	}

	return p
}

func main() {
	for _, row := range ReadInputRows() {
		fmt.Println(string(row))
		fmt.Printf("%+v\n", ParsePart(row))
	}
}
