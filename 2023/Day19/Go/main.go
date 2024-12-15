package main

import (
	"bytes"
	"fmt"
)

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

type OP bool

const (
	GT OP = true
	LT OP = false
)

type Instruction struct {
	v      byte
	op     OP
	val    int
	target string
}

func ParseInstruction(b []byte) Instruction {
	if len(b) <= 3 {
		return Instruction{
			target: string(b),
		}
	}

	resp := Instruction{
		v: b[0],
	}

	if b[1] == '>' {
		resp.op = GT
	} else {
		resp.op = LT
	}

	for i := 2; ; i++ {
		if b[i] == ':' {
			resp.val = Atoi(b[2:i])
			resp.target = string(b[i+1:])
			break
		}
	}

	return resp
}

func ParseInstructions(b []byte) []Instruction {
	insts := bytes.Split(b, []byte{','})
	resp := make([]Instruction, len(insts))

	for i, instr := range insts {
		resp[i] = ParseInstruction(instr)
	}

	return resp
}

func EvalPipline(inst []Instruction, part Part) string {
	for _, instr := range inst {
		if instr.v == 0 {
			return instr.target
		}

		switch instr.v {
		case 'x':
			if instr.op == GT && part.X > instr.val {
				return instr.target
			}

			if instr.op == LT && part.X < instr.val {
				return instr.target
			}

		case 'm':
			if instr.op == GT && part.M > instr.val {
				return instr.target
			}

			if instr.op == LT && part.M < instr.val {
				return instr.target
			}

		case 'a':
			if instr.op == GT && part.A > instr.val {
				return instr.target
			}

			if instr.op == LT && part.A < instr.val {
				return instr.target
			}

		case 's':
			if instr.op == GT && part.S > instr.val {
				return instr.target
			}

			if instr.op == LT && part.S < instr.val {
				return instr.target
			}

		}
	}

	panic("OOPS")
}

func Accepts(pipelines map[string][]Instruction, part Part) bool {
	cur := pipelines["in"]

	for {
		ans := EvalPipline(cur, part)
		if ans == "R" {
			return false
		}

		if ans == "A" {
			return true
		}

		cur = pipelines[ans]
	}
}

func ParseInput() (map[string][]Instruction, []Part) {
	input := bytes.Split(ReadInputFile(), []byte("\n\n"))
	pipeInput := bytes.Split(input[0], []byte{'\n'})

	if input[1][len(input[1])-1] == '\n' {
		input[1] = input[1][:len(input[1])-1]
	}

	partInput := bytes.Split(input[1], []byte{'\n'})

	pipelines := make(map[string][]Instruction)
	parts := make([]Part, len(partInput))

	for _, ins := range pipeInput {
		for i := 0; ; i++ {
			if ins[i] == '{' {
				pipelines[string(ins[:i])] = ParseInstructions(ins[i+1 : len(ins)-1])
				break
			}
		}
	}

	for i, row := range partInput {
		parts[i] = ParsePart(row)
	}

	return pipelines, parts
}

func PartTotal(p Part) int {
	return p.X + p.M + p.A + p.S
}

func main() {
	pipelines, parts := ParseInput()

	p1total := 0
	for _, p := range parts {
		if Accepts(pipelines, p) {
			p1total += PartTotal(p)
		}
	}

	fmt.Println("Puzzle One:", p1total)
}
