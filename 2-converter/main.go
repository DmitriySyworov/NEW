package main

import (
	"errors"
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
		original, error1, quantity, error2, final, error3, error4 := Input()
		if error1 != nil {
			fmt.Println(error1, "Таких значений для конвертации не существует")
		} else if error3 != nil {
			fmt.Println(error3, "Таких значений для конвертации не существует")
		} else if error2 != nil {
			fmt.Println(error2, "Валюта не может быть равна нулю или отрицательному значению")
		} else if error4 != nil {
			fmt.Println(error4, "Указанные валюты не могут быть одинаковы")
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
func Input() (string, error, float64, error, string, error, error) {
	var originalCurrency, finalCurrency string
	var quantity float64
	fmt.Println("Укажите вашу валюту для дальнейшей конвертации (USD/EUR/RUB)")
	fmt.Scan(&originalCurrency)
	if originalCurrency != "RUB" && originalCurrency != "EUR" && originalCurrency != "USD" {
		return "", errors.New("ERROR1"), 0, nil, "", nil, nil
	}
	fmt.Println("Укажите количество вашей валюты")
	fmt.Scan(&quantity)
	if quantity <= 0 {
		return "", nil, 0, errors.New("ERROR2"), "", nil, nil
	}
	fmt.Println("Укажите валюту в которую вы хотите конвертировать деньги (USD/EUR/RUB)")
	fmt.Scan(&finalCurrency)
	if finalCurrency != "EUR" && finalCurrency != "USD" && finalCurrency != "RUB" {
		return "", nil, 0, nil, "", errors.New("ERROR3"), nil
	}
	if originalCurrency == "EUR" && finalCurrency == "EUR" || originalCurrency == "USD" && finalCurrency == "USD" || originalCurrency == "RUB" && finalCurrency == "RUB" {
		return "", nil, 0, nil, "", nil, errors.New("ERROR4")
	}
	return originalCurrency, nil, quantity, nil, finalCurrency, nil, nil
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
