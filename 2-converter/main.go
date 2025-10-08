package main

import (
	"fmt"
)

const EURonUSD = 1.17
const EURonRUB = 97.35
const USDonEUR = 0.85
const USDonRUB = 82.9
const RUBonEUR = 0.01
const RUBonUSD = 0.012

func main() {
	fmt.Println("__Конвертор валют__")
	for {
		original := inputOriginal()
		quantity := inputQuantity()
		final := inputFinal()
		if original == final {
			fmt.Println("ERROR4: Валюты не могут быть одинаковы, пожалуйста повторите попытку")
			continue
		}
		outputCurrency(original, quantity, final)
		if choice() {
			fmt.Println("Повторный расчет:")
		} else {
			fmt.Println("Конец программы")
			break
		}
	}
}
func inputOriginal() string {
	for {
		var originalCurrency string
		fmt.Println("Укажите вашу валюту для дальнейшей конвертации (USD/EUR/RUB)")
		fmt.Scan(&originalCurrency)
		if originalCurrency != "RUB" && originalCurrency != "EUR" && originalCurrency != "USD" {
			fmt.Println("ERROR1: Таких значений для конвертации не существует, пожалуйста повторите попытку ")
			continue
		}
		return originalCurrency
	}
}
func inputQuantity() float64 {
	for {
		var quantity float64
		fmt.Println("Укажите количество вашей валюты")
		fmt.Scan(&quantity)
		if quantity <= 0 {
			fmt.Println("ERROR2: Валюта не может быть равна нулю или отрицательному значению, пожалуйста повторите попытку")
			continue
		}
		return quantity
	}
}
func inputFinal() string {
	for {
		var finalCurrency string
		fmt.Println("Укажите валюту в которую вы хотите конвертировать деньги (USD/EUR/RUB)")
		fmt.Scan(&finalCurrency)
		if finalCurrency != "EUR" && finalCurrency != "USD" && finalCurrency != "RUB" {
			fmt.Println("ERROR3: Таких значений для конвертации не существует, пожалуйста повторите попытку")
			continue
		}
		return finalCurrency
	}
}
func outputCurrency(original string, quantity float64, final string) {
	switch {
	case original == "EUR" && final == "USD":
		fmt.Println("USD = ", quantity*EURonUSD)
	case original == "EUR" && final == "RUB":
		fmt.Println("RUB = ", quantity*EURonRUB)
	case original == "USD" && final == "EUR":
		fmt.Println("EUR = ", quantity*USDonEUR)
	case original == "USD" && final == "RUB":
		fmt.Println("RUB = ", quantity*USDonRUB)
	case original == "RUB" && final == "EUR":
		fmt.Println("EUR = ", quantity*RUBonEUR)
	case original == "RUB" && final == "USD":
		fmt.Println("USD = ", quantity*RUBonUSD)
	}
}
func choice() bool {
	var choice string
	fmt.Println("Желаете ли вы повторить расчет? Укажите yes или no")
	fmt.Scan(&choice)
	if choice == "yes" || choice == "YES" || choice == "Y" || choice == "y" {
		return true
	}
	return false
}
