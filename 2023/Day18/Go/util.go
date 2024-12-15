package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
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

func HexToInt(b []byte) int {
	i, err := strconv.ParseInt(string(b), 16, 64)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func Mod(a, b int) int {
	return (a%b + b) % b
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

type Array2D [][]byte

// F* this noise
func (a Array2D) RotateCW() Array2D {
	resp := make(Array2D, len(a[0]))
	for y := 0; y < len(resp); y++ {
		resp[y] = make([]byte, len(a))
		for x := len(a) - 1; x >= 0; x-- {
			resp[y][len(a)-x-1] = a[x][y]
		}
	}

	return resp
}

func (a Array2D) RotateCCW() Array2D {
	resp := make(Array2D, len(a[0]))
	for y := 0; y < len(resp); y++ {
		resp[y] = make([]byte, len(a))
		for x := len(a) - 1; x >= 0; x-- {
			resp[y][x] = a[x][len(a[0])-y-1]
		}
	}

	return resp
}

func (a Array2D) Hash() string {
	h := md5.New()
	for _, row := range a {
		h.Write(row)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func (a Array2D) Clone() Array2D {
	resp := make(Array2D, len(a))
	for y, row := range a {
		resp[y] = make([]byte, len(row))
		for x, b := range row {
			resp[y][x] = b
		}
	}

	return resp
}
