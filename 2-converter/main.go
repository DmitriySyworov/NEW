package main

import "fmt"

func main() {
	number, eur, usd := output()
	result1, result2 := countCurrencies(number, eur, usd)
	fmt.Printf("Указанное количество рублей = %.2f EUR \n", result1)
	fmt.Printf("указанное количество рублей = %.2f USD", result2)
}
func output() (float64, float64, float64) {
	var numberCurrency, EUR, USD float64
	fmt.Println("Укажите количество рублей и через пробел курс валюты евро и доллара по отношению к рублю")
	fmt.Scan(&numberCurrency, &EUR, &USD)
	return numberCurrency, EUR, USD
}
func countCurrencies(numbers, EUR, USD float64) (float64, float64) {
	rez := numbers / EUR
	rez2 := numbers / USD
	return rez, rez2
}
