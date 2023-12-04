package main

import (
	"fmt"
	"math"
	"strconv"
)

type Card struct {
	Number  int
	Winning []int
	Have    []int
}

func (c Card) Winners() []int {
	resp := make([]int, 0, len(c.Winning))

	for _, w := range c.Winning {
		for _, h := range c.Have {
			if w == h {
				resp = append(resp, h)
			}
		}
	}

	return resp
}

func (c Card) Score() int {
	winnerCount := len(c.Winners())
	if winnerCount == 0 {
		return 0
	}

	return int(math.Pow(2, float64(winnerCount)-1))
}

func (c Card) ChildrenCards() []int {
	resp := make([]int, 0)

	for i := 1; i <= len(c.Winners()); i++ {
		resp = append(resp, c.Number+i)
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
	resp := make([]int, 0, 25)
	buff := make([]byte, 0, 2)

	for i := 0; i < len(b); i++ {
		if b[i] == ' ' && len(buff) != 0 {
			resp = append(resp, Atoi(buff))
			buff = buff[:0]

		} else if b[i] != ' ' {
			buff = append(buff, b[i])

		}
	}

	if len(buff) != 0 {
		resp = append(resp, Atoi(buff))
	}

	return resp
}

func Cards(rows [][]byte) map[int]Card {
	resp := make(map[int]Card)
	for i, row := range rows {
		resp[i+1] = Card{
			Number:  i + 1,
			Winning: SplitInts(row[10:39]),
			Have:    SplitInts(row[42:]),
			// Winning: SplitInts(row[8:23]),
			// Have:    SplitInts(row[25:]),
		}
	}

	return resp
}

func CountCards(c Card, cards map[int]Card, total *int) {
	*total++

	for _, children := range c.ChildrenCards() {
		CountCards(cards[children], cards, total)
	}
}

func main() {
	puzzleOne, puzzleTwo := 0, 0
	// puzzleOne := 0
	cards := Cards(ReadInputRows())

	for _, c := range cards {
		CountCards(c, cards, &puzzleTwo)
		puzzleOne += c.Score()
	}

	fmt.Println("Puzzle One: ", puzzleOne)
	fmt.Println("Puzzle Two: ", puzzleTwo)
}
