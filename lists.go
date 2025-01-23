package main

import "fmt"

func main() {
	var productNames [4]string
	productNames = [4]string{"A book"}
	productNames[2] = "Carpet"

	prices := [4]float64{1.99, 2.99, 3.99, 4.99}
	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[0])

	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]
	//featuredPrices := prices[:3]
	//featuredPrices := prices[1:]

	featuredPrices[0] = 199.99

	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
}
