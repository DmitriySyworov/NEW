package main

import (
	"errors"
	"fmt"
)

type currencyType = map[string]*float64

func main() {
	EU := 1.17
	ER := 97.35
	UE := 0.85
	UR := 82.9
	RE := 0.01
	RU := 0.012
	fmt.Println("Конвертор валют")
	currency := currencyType{
		"EURonUSD": &EU,
		"EURonRUB": &ER,
		"USDonEUR": &UE,
		"USDonRUB": &UR,
		"RUBonEUR": &RE,
		"RUBonUSD": &RU,
	}

	for {
		original, err1 := inputOriginal()
		if err1 != nil {
			fmt.Println(err1, "Таких значений для конвертации не существует, пожалуйста повторите попытку")
			continue
		}
		quantity, err2 := inputQuantity()
		if err2 != nil {
			fmt.Println(err2, "Валюта не может быть равна нулю или отрицательному значению, пожалуйста повторите попытку")
			continue
		}
		final, err3 := inputFinal()
		if err3 != nil {
			fmt.Println(err3, "Таких значений для конвертации не существует, пожалуйста повторите попытку")
			continue
		}
		if original == final {
			fmt.Println("ERROR4: Валюты не могут быть одинаковы, пожалуйста повторите попытку")
			continue
		}
		fmt.Println(coutCurrency(currency[original+"on"+final], quantity))
		if choice() {
			fmt.Println("Повторный расчет:")
		} else {
			fmt.Println("Конец программы")
			break
		}
	}
}
func inputOriginal() (string, error) {
	var originalCurrency string
	fmt.Println("Укажите вашу валюту для дальнейшей конвертации (USD/EUR/RUB)")
	fmt.Scan(&originalCurrency)
	if originalCurrency != "RUB" && originalCurrency != "EUR" && originalCurrency != "USD" {
		return "", errors.New("ERROR1:")
	}
	return originalCurrency, nil
}

func inputQuantity() (float64, error) {
	var quantity float64
	fmt.Println("Укажите количество вашей валюты")
	fmt.Scan(&quantity)
	if quantity <= 0 {
		return 0, errors.New("ERROR2:")
	}
	return quantity, nil
}
func inputFinal() (string, error) {
	var finalCurrency string
	fmt.Println("Укажите валюту в которую вы хотите конвертировать деньги (USD/EUR/RUB)")
	fmt.Scan(&finalCurrency)
	if finalCurrency != "EUR" && finalCurrency != "USD" && finalCurrency != "RUB" {
		return "", errors.New("ERROR3:")
	}
	return finalCurrency, nil
}
func coutCurrency(maps *float64, quantity float64) float64 {
	return *maps * quantity
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
