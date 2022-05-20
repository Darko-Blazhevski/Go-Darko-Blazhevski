package main

import (
	"fmt"

	"Lecture8-1/carddraw"
	"Lecture8-1/cardgame"
)

func main() {

	aaap := cardgame.NewDeck()
	// b:=&aaap
	// fmt.Println(aaap.Deal())
	// fmt.Println(aaap.Deal())
	// fmt.Println(aaap.Deal())
	// fmt.Println(aaap.Deal())
	// fmt.Println(aaap.Deal())
	// fmt.Println(aaap.Deal())
	// fmt.Println(aaap.Deal())
	// fmt.Println(aaap.Deal())
	// fmt.Println(aaap.Deal())
	// fmt.Println(aaap.Deal())
	fmt.Println(carddraw.DrawAllCards(&aaap))

	// fmt.Println(aaap)
	// fmt.Println(carddraw.DrawAllCards(aaap))

}

func maxCard(cards []cardgame.Card, comparatorFunc CardComparator) cardgame.Card {
	maximumCard := cardgame.NewCard("hearts", 2)
	for i := 0; i < len(cards); i++ {
		if comparatorFunc.compareCards(maximumCard, cards[i]) == -1 {
			maximumCard = cards[i]
		}
	}
	return maximumCard
}

type CardComparator func(cardOne cardgame.Card, cardTwo cardgame.Card) int

func (comparatorFunc CardComparator) compareCards(cardOne cardgame.Card, cardTwo cardgame.Card) int {
	if cardOne.Suit != "none" || cardOne.Strenght != 0 || cardTwo.Suit != "none" || cardTwo.Strenght != 0 {
		if cardOne.Strenght > cardTwo.Strenght {
			return 1
		} else if cardOne.Strenght < cardTwo.Strenght {
			return -1
		} else {
			if cardOne.Suit == "diamonds" && cardTwo.Suit == "diamonds" {
				return 0
			} else if cardOne.Suit == "diamonds" && cardTwo.Suit == "clubs" {
				return 1
			} else if cardOne.Suit == "diamonds" && cardTwo.Suit == "spades" {
				return 1
			} else if cardOne.Suit == "diamonds" && cardTwo.Suit == "hearts" {
				return 1
			} else if cardOne.Suit == "clubs" && cardTwo.Suit == "diamonds" {
				return -1
			} else if cardOne.Suit == "clubs" && cardTwo.Suit == "clubs" {
				return 0
			} else if cardOne.Suit == "clubs" && cardTwo.Suit == "spades" {
				return 1
			} else if cardOne.Suit == "clubs" && cardTwo.Suit == "hearts" {
				return 1
			} else if cardOne.Suit == "spades" && cardTwo.Suit == "diamonds" {
				return -1
			} else if cardOne.Suit == "spades" && cardTwo.Suit == "clubs" {
				return -1
			} else if cardOne.Suit == "spades" && cardTwo.Suit == "spades" {
				return 0
			} else if cardOne.Suit == "spades" && cardTwo.Suit == "hearts" {
				return 1
			} else if cardOne.Suit == "hearts" && cardTwo.Suit == "diamonds" {
				return -1
			} else if cardOne.Suit == "hearts" && cardTwo.Suit == "clubs" {
				return -1
			} else if cardOne.Suit == "hearts" && cardTwo.Suit == "spades" {
				return -1
			} else if cardOne.Suit == "hearts" && cardTwo.Suit == "hearts" {
				return 0
			}
		}
	} else {
		fmt.Println("Invalid card parametars")
		return 0
	}
	return 0
}
