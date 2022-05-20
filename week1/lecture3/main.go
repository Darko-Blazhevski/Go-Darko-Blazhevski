package main

import "fmt"

func main() {

	fmt.Println("Comapering cards!")

	fmt.Println(compareCards(10, 5, 10, 2))

}

func compareCards(cardOneVal int, cardOneSuit int, cardTwoVal int, cardTwoSuit int) int {
	if cardOneVal >= 2 && cardOneVal <= 14 && cardTwoVal >= 2 && cardTwoVal <= 14 {
		if cardOneSuit >= 1 && cardOneSuit <= 4 && cardTwoSuit >= 1 && cardTwoSuit <= 4 {
			if cardOneVal > cardTwoVal {
				return 1
			} else if cardTwoVal > cardOneVal {
				return -1
			} else if cardTwoVal == cardOneVal {
				if cardOneSuit > cardTwoSuit {
					return 1
				} else if cardOneSuit > cardTwoSuit {
					return -1
				} else if cardOneSuit == cardTwoSuit {
					return 0
				}

			}

		} else {
			fmt.Println("Card suit is incorrect!")
			return 0
		}
	} else {
		fmt.Println("Card Value is incorrect!")
		return 0
	}
	return 0
}
