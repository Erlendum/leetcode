package main

import "fmt"

func maxProfit(prices []int) int {
	mProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			mProfit += prices[i] - prices[i-1]
		}
	}
	return mProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Print(maxProfit(prices))
}
