package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := getArray()
	fmt.Println(prices)

	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	startIndex := 1
	featuredPrices := prices[startIndex:]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)

	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	fmt.Println(highlightedPrices[2:cap(highlightedPrices)])
}

func getArray() []float64 {
	return []float64{10.99, 9.99, 45.99, 20.0}
}
