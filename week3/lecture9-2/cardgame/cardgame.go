package cardgame

import "errors"

// import "math/rand"

type Card struct {
	Suit     string
	Strenght int
}

func NewCard(Suit string, Strenght int) Card {
	if Suit == "diamonds" || Suit == "clubs" || Suit == "spades" || Suit == "hearts" {
		if Strenght <= 14 && Strenght >= 2 {
			return Card{Suit: Suit, Strenght: Strenght}

		} else {
			return Card{Suit: "none", Strenght: 0}
		}
	} else {
		return Card{Suit: "none", Strenght: 0}
	}

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
// func (d *Deck) Shuffle() {
// 	rand.Shuffle(len(d.spil), func(i, j int) {
// 		d.spil[i], d.spil[j] = d.spil[j], d.spil[i]
// 	})
// }
func (d *Deck) Deal() (Card ,error) {
	var karta Card
	if len(d.spil) == 0 {
		return Card{Suit: "none",Strenght: 0}, errors.New("Deck is empty")
	} else {
		karta=d.spil[len(d.spil)-1]
		d.spil = d.spil[:len(d.spil)-1]
		return karta, nil
		
	}
}
func (d *Deck) Done() bool  {
	if len(d.spil)==0 {
		return true
	}else {
		return false
	}
}
