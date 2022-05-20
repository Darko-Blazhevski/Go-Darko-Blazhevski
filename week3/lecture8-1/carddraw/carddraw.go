package carddraw

import "Lecture8-1/cardgame"
// import "fmt"

type dealer interface {
	Deal() cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {
	spilkarti:= []cardgame.Card{}
	for {
		val:=dealer.Deal()
		
		if val.Strenght==0 {
			break
		}else {
			
			spilkarti = append(spilkarti, val)
		}
	}
 return spilkarti
}
// func DrawAllCards(dealer dealer) []cardgame.Card {
// 	spilkarti:= []cardgame.Card{}
// 	for val:=dealer.Deal();val.Strenght!=0; {
// 			spilkarti = append(spilkarti,val)
// 	}
//  return spilkarti
// }

// func DrawAllCards(dealer dealer) []cardgame.Card {
// 	spilkarti:= []cardgame.Card{}
// 	for val:=dealer.Deal();val.Strenght!=0; {
// 		if val.Strenght!=0 {
			
// 			fmt.Println(val)
// 			spilkarti = append(spilkarti,dealer.Deal())
// 			// fmt.Println(spilkarti)
// 		}else {
// 			fmt.Println("kraj")
// 			break
// 		}
// 	}
//  return spilkarti
// }

// func DrawAllCards(dealer dealer) []cardgame.Card {
// 	var spilkarti []cardgame.Card
// 	for val:=dealer.Deal();val!=nil;{
// 		spilkarti=append(spilkarti,*val)
// 	}
// 	return spilkarti
// }

// func DrawAllCards(dealer dealer) []cardgame.Card {
// 	spilkarti:= []cardgame.Card{}
// 	for {
// 		if dealer.Deal().Strenght==0 {
// 			break
// 		}else {
			
// 			spilkarti = append(spilkarti, cardgame.Card{Suit: dealer.Deal().Suit,Strenght: dealer.Deal().Strenght})
// 		}
// 	}
//  return spilkarti
// }
