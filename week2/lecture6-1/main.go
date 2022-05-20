package main

import "fmt"

func main() {

	// p1:=Item{Value: 10}
	// p2:=Item{Value: 12}
	// p3:=Item{Value: 14}
	// p4:=Item{Value: 16}

	// //fmt.Println(p1)

	// l1:=&MagicList{&p1}
	// l2:=&MagicList{&p2}
	// l3:=&MagicList{&p3}
	// l4:=&MagicList{&p4}

	// s1:=[]MagicList{*l1,*l2,*l3,*l4}

	// //fmt.Println(l1,"\n",l2)

	// fmt.Println(s1)

	lala := MagicList{}

	fmt.Println(lala)

	add(&lala, 15)
	fmt.Println(lala)
	add(&lala, 17)
	fmt.Println(lala)
	add(&lala, 8)
	fmt.Println(lala)
	add(&lala, 14)
	fmt.Println(lala)
	add(&lala, 15)
	fmt.Println(lala)

	fmt.Println(toSlice(&lala))

}

type MagicList struct {
	LastItem *Item
}

type Item struct {
	Value    int
	PrevItem *Item
}

func add(l *MagicList, value int) {
	ni := Item{Value: value}
	if l.LastItem == nil {
		l.LastItem = &ni
	} else if l.LastItem != nil {
		ni.PrevItem = l.LastItem
		l.LastItem = &ni
	}

}

func toSlice(l *MagicList) []int {
	swice := []int{}
	piece := l.LastItem
	for {
		if piece == nil {
			break
		}
		swice = append(swice, piece.Value)
		piece = piece.PrevItem
	}
	return swice
}
