package main

import (
	"bytes"
	"os"
)

func ReadInputFile() []byte {
	if len(os.Args) < 2 {
		panic("Missing input file arg")
	}

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	return b
}

func ReadInputRows() [][]byte {
	b := bytes.Split(ReadInputFile(), []byte{'\n'})
	if len(b[len(b)-1]) == 0 {
		return b[:len(b)-1]
	}

	return b
}
