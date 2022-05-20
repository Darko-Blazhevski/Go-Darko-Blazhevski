package main

import "fmt"

func main() {

	fmt.Println(daysInMonth(2, 1924))

}

func daysInMonth(month int, year int) (int, bool) {
	if month >= 1 && month <= 12 {
		if year > 0 {
			switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				return 31, true
			case 4, 6, 9, 11:
				return 30, true
			case 2:
				if year%4 == 0 {
					return 29, true
				} else {
					return 28, true
				}
			}
		} else {
			return 0, false
		}
	} else {
		return 0, false
	}
	return 0, false
}
