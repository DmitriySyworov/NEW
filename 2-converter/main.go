package main

import (
	"errors"
	"fmt"
)

type currencyType = map[string]float64

func main() {
	fmt.Println("__Конвертор валют__")
	variableCurrency := currencyType{
		"EURonUSD": 1.17,
		"EURonRUB": 97.35,
		"USDonEUR": 0.85,
		"USDonRUB": 82.9,
		"RUBonEUR": 0.01,
		"RUBonUSD": 0.012,
	}
	for {
		quantity, err1 := inputAndChoiceCurrencyTransaction()

		if err1 != nil {
			fmt.Println(err1, "Валюта не может быть равна нулю или отрицательному значению, пожалуйста повторите попытку")
			continue
		}
		outputCurrency(quantity, variableCurrency)
		if choice() {
			fmt.Println("Повторный расчет:")
		} else {
			fmt.Println("Конец программы")
			break
		}
	}
}
func inputAndChoiceCurrencyTransaction() (float64, error) {
	fmt.Println(`Укажите количество валюты для ее рассчета, где первая валюта - ваша изначальная:
EURonUSD - из евро в доллары
EURonRUB -  из евро в рубли
USDonEUR - из долларов в евро
USDonRUB - из долларов в рубли
RUBonEUR - из рублей в евро
RUBonUSD - из рублей в доллары`)
	var quantity float64
	fmt.Scan(&quantity)
	if quantity <= 0 {
		return 0, errors.New("ERROR2:")
	}
	return quantity, nil
}

func outputCurrency(quantity float64, variableCurrency currencyType) {
	variableCurrency2 := currencyType{}
	for key, value := range variableCurrency {
		variableCurrency2[key] = value * quantity
		fmt.Printf("%s = %.2f \n", key, variableCurrency2[key])
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
