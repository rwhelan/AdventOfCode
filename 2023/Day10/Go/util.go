package main

import (
	"bytes"
	"os"
	"strconv"
)

type Map2D map[int]map[int]byte

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

func ReadInputMap() Map2D {
	resp := make(Map2D)
	for y, row := range ReadInputRows() {
		resp[y] = make(map[int]byte)
		for x, b := range row {
			resp[y][x] = b
		}
	}

	return resp
}

func Atoi(b []byte) int {
	num, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}

	return num
}

func SplitInts(b []byte) []int {
	resp := make([]int, 0, 64)
	buff := make([]byte, 0)

	for i, c := range b {
		if IsDigit(c) || c == '-' {
			buff = append(buff, c)
		}

		if len(buff) != 0 && c == ' ' || i == len(b)-1 {
			resp = append(resp, Atoi(buff))
			buff = buff[:0]
		}
	}

	return resp
}

func IsDigit(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	}

	return false
}
