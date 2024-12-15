package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

var cardOrder = map[byte]int{
	'2': 1, '3': 2, '4': 3, '5': 4,
	'6': 5, '7': 6, '8': 7, '9': 8,
	'T': 9, 'J': 10, 'Q': 11, 'K': 12,
	'A': 13,
}

type Hands []Hand

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hands) Less(i, j int) bool {
	return h[i].Beats(h[j])
}

type Hand struct {
	cards []byte
	bid   int
}

func handValues(cards []byte) []int {
	m := make(map[byte]int)
	for _, c := range cards {
		m[c]++
	}

	n := make([]int, len(cards))
	for i, c := range cards {
		n[i] = m[c]
	}

	slices.Sort(n)
	slices.Reverse(n)

	return n
}

func handNumber(cards []byte) int {
	n := handValues(cards)

	slices.Reverse(n)
	resp := 0
	for i := 0; i < len(cards); i++ {
		resp += n[i] * int(math.Pow10(i))
	}

	return resp
}

func addjoker(values []int) []int {
	resp := append([]int{values[0]}, values...)
	for i := 0; i < resp[0]; i++ {
		resp[i]++
	}

	return resp
}

func handNumberWithJokers(cards []byte) int {
	jokers := 0
	for _, c := range cards {
		if c == 'J' {
			jokers++
		}
	}

	if jokers == 5 {
		return 55555
	}

	cardsWOjokers := make([]byte, 0, len(cards)-jokers)
	for _, c := range cards {
		if c != 'J' {
			cardsWOjokers = append(cardsWOjokers, c)
		}
	}

	n := handValues(cardsWOjokers)
	for i := 0; i < jokers; i++ {
		n = addjoker(n)
	}

	slices.Reverse(n)
	resp := 0
	for i := 0; i < len(cards); i++ {
		resp += n[i] * int(math.Pow10(i))
	}

	return resp
}

func (h Hand) Beats(other Hand) bool {
	thisHandType := handNumberWithJokers(h.cards)
	otherHandType := handNumberWithJokers(other.cards)

	if thisHandType != otherHandType {
		return thisHandType > otherHandType
	}

	for i := 0; i < len(h.cards); i++ {
		if cardOrder[h.cards[i]] == cardOrder[other.cards[i]] {
			continue
		}

		return cardOrder[h.cards[i]] > cardOrder[other.cards[i]]
	}

	panic("CARDS ARE THE SAME")
}

func Parse() Hands {
	rows := ReadInputRows()
	resp := make(Hands, len(rows))

	for i, row := range rows {
		resp[i] = Hand{
			cards: row[:5],
			bid:   Atoi(row[6:]),
		}
	}

	return resp
}

func main() {
	t := 0
	hands := Parse()

	sort.Sort(hands)
	slices.Reverse(hands)

	for i, h := range hands {
		t += (i + 1) * h.bid
	}

	fmt.Println("Puzzle One:", t)
}
