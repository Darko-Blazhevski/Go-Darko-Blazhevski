package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// k1 := NewCard("spades", 13)
	// k2 := NewCard("diamonds", 3)
	// k3 := NewCard("hearts", 9)
	// k4 := NewCard("spades", 8)
	// k5 := NewCard("clubs", 13)
	// fmt.Println(compareCards(k2, k3))
	// slice := []Card{k1, k2, k3, k4, k5}
	// slice=append(slice,NewCard("spades", 13))
	// fmt.Println(maxCard(slice))
	// fmt.Println(slice)
	aaap := NewDeck()
	// fmt.Println(aaap)
	// fmt.Println(len(aaap.spil))

	// fmt.Println(aaap.Shuffle(aaap))
	aaap.Shuffle()
	// fmt.Println(aaap)
	aaap.Deal()

	var neznampoke CardComparator

	// neznampoke(k1,k2)
	// fmt.Println(aaap)
	fmt.Println(maxCard(aaap.spil, neznampoke))

}

type Deck struct {
	spil []Card
}

func NewDeck() Deck {
	spil := []Card{}
	for i := 2; i < 15; i++ {
		spil = append(spil, NewCard("diamonds", i))
	}
	for i := 2; i < 15; i++ {
		spil = append(spil, NewCard("clubs", i))
	}
	for i := 2; i < 15; i++ {
		spil = append(spil, NewCard("spades", i))
	}
	for i := 2; i < 15; i++ {
		spil = append(spil, NewCard("hearts", i))
	}

	return Deck{spil: spil}
}
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.spil), func(i, j int) {
		d.spil[i], d.spil[j] = d.spil[j], d.spil[i]
	})
}
func (d *Deck) Deal() {
	if len(d.spil) == 0 {
		fmt.Println("Deck is empty")
	} else {
		fmt.Println(d.spil[len(d.spil)-1])
		d.spil = d.spil[:len(d.spil)-1]
	}

}

// type CardComparator func(cOne Card, cTwo Card) int

func maxCard(cards []Card, comparatorFunc CardComparator) Card {
	maximumCard := NewCard("hearts", 2)
	for i := 0; i < len(cards); i++ {
		if comparatorFunc.compareCards(maximumCard, cards[i]) == -1 {
			maximumCard = cards[i]
		}
	}
	return maximumCard
}

// func (Deck)Shuffle(lips Deck) Deck {
// 	rand.Shuffle(len(lips.spil), func(i,j int)  {
// 		lips.spil[i],lips.spil[j]=lips.spil[j],lips.spil[i]
// 	})
// 	return lips
// }
type Card struct {
	suit     string
	strenght int
}

func NewCard(suit string, strenght int) Card {
	if suit == "diamonds" || suit == "clubs" || suit == "spades" || suit == "hearts" {
		if strenght <= 14 && strenght >= 2 {
			return Card{suit: suit, strenght: strenght}

		} else {
			return Card{suit: "none", strenght: 0}
		}
	} else {
		return Card{suit: "none", strenght: 0}
	}

}

type CardComparator func(cardOne Card, cardTwo Card) int

func (comparatorFunc CardComparator) compareCards(cardOne Card, cardTwo Card) int {
	if cardOne.suit != "none" || cardOne.strenght != 0 || cardTwo.suit != "none" || cardTwo.strenght != 0 {
		if cardOne.strenght > cardTwo.strenght {
			return 1
		} else if cardOne.strenght < cardTwo.strenght {
			return -1
		} else {
			if cardOne.suit == "diamonds" && cardTwo.suit == "diamonds" {
				return 0
			} else if cardOne.suit == "diamonds" && cardTwo.suit == "clubs" {
				return 1
			} else if cardOne.suit == "diamonds" && cardTwo.suit == "spades" {
				return 1
			} else if cardOne.suit == "diamonds" && cardTwo.suit == "hearts" {
				return 1
			} else if cardOne.suit == "clubs" && cardTwo.suit == "diamonds" {
				return -1
			} else if cardOne.suit == "clubs" && cardTwo.suit == "clubs" {
				return 0
			} else if cardOne.suit == "clubs" && cardTwo.suit == "spades" {
				return 1
			} else if cardOne.suit == "clubs" && cardTwo.suit == "hearts" {
				return 1
			} else if cardOne.suit == "spades" && cardTwo.suit == "diamonds" {
				return -1
			} else if cardOne.suit == "spades" && cardTwo.suit == "clubs" {
				return -1
			} else if cardOne.suit == "spades" && cardTwo.suit == "spades" {
				return 0
			} else if cardOne.suit == "spades" && cardTwo.suit == "hearts" {
				return 1
			} else if cardOne.suit == "hearts" && cardTwo.suit == "diamonds" {
				return -1
			} else if cardOne.suit == "hearts" && cardTwo.suit == "clubs" {
				return -1
			} else if cardOne.suit == "hearts" && cardTwo.suit == "spades" {
				return -1
			} else if cardOne.suit == "hearts" && cardTwo.suit == "hearts" {
				return 0
			}
		}
	} else {
		fmt.Println("Invalid card parametars")
		return 0
	}
	return 0
}

// func maxCard(cards []Card) Card {
// 	maximumCard := NewCard("hearts", 2)
// 	for i := 0; i < len(cards); i++ {
// 		if compareCards(maximumCard, cards[i]) == -1 {
// 			maximumCard = cards[i]
// 		}
// 	}
// 	return maximumCard
// }
