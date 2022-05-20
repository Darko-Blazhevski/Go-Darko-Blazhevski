package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	a, b := citiesAndPrices()

	m := groupSlices(a, b)

	fmt.Println(m)

}

func groupSlices(cities []string, Prices []int) map[string][]int {
	var mapa map[string][]int
	for i := 0; i < len(cities); i++ {
		if mapa == nil {
			mapa = make(map[string][]int)
		} else {
			mapa[cities[i]] = append(mapa[cities[i]], Prices[i])
		}

	}
	return mapa
}

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}
