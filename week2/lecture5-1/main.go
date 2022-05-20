package main

import "fmt"

func main() {

	k1, _ := NewCard("spades", 13)
	k2, _ := NewCard("diamonds", 3)
	k3, _ := NewCard("hearts", 9)
	k4, _ := NewCard("spades", 8)
	k5, _ := NewCard("clubs", 13)
	fmt.Println(compareCards(k2, k3))
	slice := []Card{k1, k2, k3, k4, k5}
	fmt.Println(maxCard(slice))
}

type Card struct {
	suit     string
	strenght int
}

func NewCard(suit string, strenght int) (Card, bool) {
	if suit == "diamonds" || suit == "clubs " || suit == "spades" || suit == "hearts " {
		if strenght <= 14 && strenght >= 2 {
			return Card{suit: suit, strenght: strenght}, true

		} else {
			return Card{suit: "none", strenght: 0}, false
		}
	} else {
		return Card{suit: "none", strenght: 0}, false
	}

}
func compareCards(cardOne Card, cardTwo Card) int {
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

func maxCard(cards []Card) Card {
	maximumCard, _ := NewCard("hearts", 2)
	for i := 0; i < len(cards); i++ {
		if compareCards(maximumCard, cards[i]) == -1 {
			maximumCard = cards[i]
		}
	}
	return maximumCard
}
