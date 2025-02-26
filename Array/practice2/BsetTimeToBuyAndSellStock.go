package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	maxSell := 0

	minPrice := prices[0]

	for num := range prices {
		if minPrice > num {
			minPrice = num
		} else {
			if maxSell < num-minPrice {
				maxSell = num - minPrice
			}
		}
	}
	return maxSell
}
