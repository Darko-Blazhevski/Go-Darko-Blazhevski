package carddraw

import "Lecture9-2/cardgame"
// import "fmt"

type dealer interface {
	Deal() (cardgame.Card, error)
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
	spilkarti:= []cardgame.Card{}
	for {
		
		
		if !dealer.Done() {
			val,_:=dealer.Deal()
			spilkarti = append(spilkarti, val)
			
		}else {
			
			return spilkarti, nil
		}
	}
}

// func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
// 	spilkarti:= []cardgame.Card{}
// 	for {
// 		val,_:=dealer.Deal()
		
// 		if val.Strenght==0 {
// 			break
// 		}else {
			
// 			spilkarti = append(spilkarti, val)
// 		}
// 	}
//  return spilkarti, nil
// }
