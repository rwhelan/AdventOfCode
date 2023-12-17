package main

import (
	"bytes"
	"fmt"
)

type Lens struct {
	FocalLength int
	Label       string
	Next        *Lens
}

func Hash(b []byte) int {
	resp := 0

	for _, c := range b {
		resp += int(c)
		resp *= 17
		resp %= 256
	}

	return resp
}

func PuzzleOne(input []byte) int {
	t := 0
	for _, seq := range bytes.Split(input, []byte(",")) {
		t += Hash(seq)
	}

	return t
}

func InsertLens(boxes []*Lens, label []byte, fl int) {
	boxId := Hash(label)

	if boxes[boxId] == nil {
		boxes[boxId] = &Lens{
			Label:       string(label),
			FocalLength: fl,
		}

		return
	}

	c := boxes[boxId]
	var prev *Lens

	for c != nil {
		if c.Label == string(label) {
			c.FocalLength = fl
			return
		}

		prev = c
		c = c.Next
	}

	prev.Next = &Lens{
		Label:       string(label),
		FocalLength: fl,
	}
}

func RemoveLens(boxes []*Lens, label []byte) {
	boxId := Hash(label)

	if boxes[boxId] == nil {
		return
	}

	c := boxes[boxId]

	if c.Label == string(label) {
		if c.Next != nil {
			boxes[boxId] = c.Next
		} else {
			boxes[boxId] = nil
		}
		return
	}

	var prev *Lens
	for c != nil {
		if c.Label == string(label) {
			prev.Next = c.Next
		}
		prev = c
		c = c.Next
	}
}

func PuzzleTwo(input []byte) int {
	boxes := make([]*Lens, 256)

	for _, seq := range bytes.Split(input, []byte(",")) {
		switch {
		case seq[len(seq)-1] == '-':
			RemoveLens(boxes, seq[:len(seq)-1])
		case seq[len(seq)-2] == '=':
			InsertLens(boxes, seq[:len(seq)-2], int(seq[len(seq)-1])-0x30)
		default:
			fmt.Println("Unsure what to do with:", string(seq))
		}
	}

	t := 0
	for i, l := range boxes {
		c := 0
		for l != nil {
			c++
			t += (i + 1) * c * l.FocalLength
			l = l.Next
		}
	}

	return t
}

func main() {
	input := ReadInputFile()
	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}

	fmt.Println("Puzzle One:", PuzzleOne(input))
	fmt.Println("Puzzle Two:", PuzzleTwo(input))
}
