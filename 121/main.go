package main

import "fmt"

func myMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func maxProfit(prices []int) int {
	l := 0
	r := 0
	mProfit := 0
	for r < len(prices) {
		curProfit := prices[r] - prices[l]
		if prices[l] < prices[r] {
			mProfit = myMax(mProfit, curProfit)
		} else {
			l = r
		}
		r++
	}
	return mProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Print(maxProfit(prices))
}
