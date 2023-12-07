package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

// type HAND_TYPE int

// const (
// 	FIVE_OF_A_KIND HAND_TYPE = iota
// 	FOUR_OF_A_KIND
// 	FULL_HOUSE
// 	THREE_OF_A_KIND
// 	TWO_PAIR
// 	ONE_PAIR
// 	HIGH_CARD
// )

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

// Converts hand into number (type)
// QQQJJ = 33322
// QQQJA = 33311
func handNumberValue(h Hand) int {
	m := make(map[byte]int)
	for _, c := range h.cards {
		m[c]++
	}

	n := make([]int, len(h.cards))
	for i, c := range h.cards {
		n[i] = m[c]
	}

	slices.Sort(n)

	resp := 0
	for i := 0; i < len(h.cards); i++ {
		resp += n[i] * int(math.Pow10(i))
	}

	return resp
}

func (h Hand) Beats(other Hand) bool {
	thisHandType := handNumberValue(h)
	otherHandType := handNumberValue(other)

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
