package main

import "strconv"

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
		if IsDigit(c) {
			buff = append(buff, c)
		}

		if c == ' ' || i == len(b)-1 {
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
