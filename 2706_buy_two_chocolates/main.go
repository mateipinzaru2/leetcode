package main

import (
	"fmt"
	"sort"
	"strconv"
)

func buyTwoChocolates(prices []float64, money float64) float64 {
	sort.Float64s(prices)

	cheapest := prices[0]
	secondCheapest := prices[1]

	if cheapest + secondCheapest > money {
		return money
	}
	return money - (cheapest + secondCheapest)
}

func main() {
	var priceInput string
	prices := []float64{}
	for {
		fmt.Print("Enter chocolate prices (or 'done' to finish): ")
		fmt.Scanln(&priceInput)
		if priceInput == "done" {
			break
		}
		price, err := strconv.ParseFloat(priceInput, 64)
		if err != nil || price <= 0 {
			fmt.Println("Invalid price:", priceInput)
			continue
		}
		prices = append(prices, price)
	}

	var moneyInput string
	fmt.Print("Enter how much money you have: ")
	fmt.Scanln(&moneyInput)
	money, err := strconv.ParseFloat(moneyInput, 64)
	if err != nil {
		fmt.Println("Invalid money:", moneyInput)
		return
	}

	leftover := buyTwoChocolates(prices, money)
	fmt.Printf("After buying the two cheapest chocolates, your leftover money is: %.1f\n", leftover)
}
