package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Matrix =====================================================================

type Matrix map[int]map[int]map[int]*Block

func (m Matrix) AddBlock(b *Block) {
	for _, c := range b.coords {
		if m[c.Z][c.Y][c.X] != nil {
			fmt.Println("Can not add block over currently existing block")
			fmt.Printf("%#v\n", *m[c.Z][c.Y][c.X])
			panic("Block exists")
		}

		m[c.Z][c.Y][c.X] = b
	}

	b.matrix = m
}

func (m Matrix) GetCoord(c Coord) *Block {
	return m[c.Z][c.Y][c.X]
}

func NewMatrix(h int) Matrix {
	resp := make(Matrix)
	for z := 0; z < h; z++ {
		resp[z] = make(map[int]map[int]*Block)
		for y := 0; y < 10; y++ {
			resp[z][y] = make(map[int]*Block)
		}
	}

	return resp
}

// Coords =====================================================================

type Coord struct {
	X, Y, Z int
}

func (c Coord) Below() Coord {
	return Coord{
		X: c.X,
		Y: c.Y,
		Z: c.Z - 1,
	}
}

func (c Coord) Above() Coord {
	return Coord{
		X: c.X,
		Y: c.Y,
		Z: c.Z + 1,
	}
}

// Block ======================================================================

type Block struct {
	Label  string
	matrix Matrix
	coords []*Coord
	o      string
}

func (b Block) belowBlocks() []*Block {
	resp := make([]*Block, 0, 8)
	for _, c := range b.coords {
		if block := b.matrix.GetCoord(c.Below()); block != nil {
			resp = append(resp, block)
		}
	}

	return resp
}

func (b Block) aboveBlocks() []*Block {
	resp := make([]*Block, 0, 8)
	for _, c := range b.coords {
		if block := b.matrix.GetCoord(c.Above()); block != nil {
			resp = append(resp, block)
		}
	}

	return resp
}

func (b *Block) Descend() bool {
	descended := false
	for ; len(b.belowBlocks()) == 0 && !b.IsOnGround(); b.lower(1) {
		descended = true
	}

	return descended
}

func (b *Block) IsOnGround() bool {
	for _, c := range b.coords {
		if c.Z == 1 {
			return true
		}
	}

	return false
}

func (b *Block) lower(count int) {
	for _, c := range b.coords {
		if c.Z == 1 {
			fmt.Println(b.Label, "Already at bottom")
			panic("Can't Descend; at ground")
		}
		b.matrix[c.Z][c.Y][c.X] = nil
		c.Z -= count
		b.matrix[c.Z][c.Y][c.X] = b
	}
}

func ParseBlock(row []byte, label string) *Block {
	sep := bytes.Index(row, []byte{'~'})
	if sep == -1 {
		fmt.Println("Bad Input Line")
		panic(string(row))
	}

	x0, y0, z0 := Atoi(row[0:1]), Atoi(row[2:3]), Atoi(row[4:sep])
	x1, y1, z1 := Atoi(row[sep+1:sep+2]), Atoi(row[sep+3:sep+4]), Atoi(row[sep+5:])
	coords := make([]*Coord, 0)

	if x0 != x1 {
		for x := x0; x <= x1; x++ {
			coords = append(coords, &Coord{X: x, Y: y0, Z: z0})
		}
	}

	if y0 != y1 {
		for y := y0; y <= y1; y++ {
			coords = append(coords, &Coord{X: x0, Y: y, Z: z0})
		}
	}

	if z0 != z1 {
		for z := z0; z <= z1; z++ {
			coords = append(coords, &Coord{X: x0, Y: y0, Z: z})
		}
	}

	return &Block{
		Label:  label,
		coords: coords,
		o:      string(row),
	}
}

// Funcs ======================================================================

func main() {
	height := 375
	matrix := NewMatrix(height)
	blocks := make([]*Block, 0)

	for _, row := range ReadInputRows() {
		h := md5.New()
		h.Write(row)
		label := hex.EncodeToString(h.Sum(nil))

		block := ParseBlock(row, label)
		matrix.AddBlock(block)

		blocks = append(blocks, block)
	}

	for z := 0; z < height; z++ {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if b := matrix.GetCoord(Coord{X: x, Y: y, Z: z}); b != nil {
					b.Descend()
				}
			}
		}
	}

	count := 0
	for i, b := range blocks {
		if len(b.aboveBlocks()) == 0 {
			fmt.Println("DO", i)
			count++
		} else {
			fmt.Println(i, b.o)
		}
	}

	fmt.Println(count)

}
