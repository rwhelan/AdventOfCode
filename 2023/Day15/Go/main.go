package main

import (
	"bytes"
	"fmt"
)

func Hash(b []byte) int {
	resp := 0

	for _, c := range b {
		resp += int(c)
		resp *= 17
		resp %= 256
	}

	return resp
}

func main() {
	input := ReadInputFile()
	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}

	t := 0
	for _, seq := range bytes.Split(input, []byte(",")) {
		fmt.Println(string(seq), Hash(seq))
		t += Hash(seq)
	}
	fmt.Println(t)
}
