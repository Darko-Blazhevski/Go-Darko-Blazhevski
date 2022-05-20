package main

import "fmt"

func main() {
	inputs := []int{1, 17, 34, 56, 2, 8}

	evenCh := processEven(inputs)
	oddCh := processEven(inputs)

	for v := range evenCh {
		fmt.Println("Even", v)
	}
	for v := range oddCh {
		fmt.Println("Odd", v)
	}

}

func processEven(inputs []int) chan int {

	che := make(chan int)

	go func() {
		for i := 0; i < len(inputs); i++ {
			if inputs[i]%2 == 0 {
				che <- inputs[i]

			}
		}
		close(che)
	}()

	return che
}
func processOdd(inputs []int) chan int {
	che := make(chan int)

	go func() {
		for i := 0; i < len(inputs); i++ {
			if inputs[i]%2 == 1 {
				che <- inputs[i]

			}
		}
		close(che)
	}()

	return che
}
