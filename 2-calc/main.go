package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("__Калькулятор транзакций__")
	for {
		choice, transaction, err, err2 := inputAndConversion()
		if err != nil {
			fmt.Println(err, "Такого вариант выбора не существует, пожалуйста повторите ввод")
			continue
		}
		if err2 != nil {
			fmt.Println(err2, "С одной транзакцией невозможно осуществлять операции, повторите ввод")
			continue
		}
		fmt.Println(Count(choice, transaction))
		repeatProgram := choiceUsers()
		if repeatProgram {
		} else {
			break
		}
	}
}
func inputAndConversion() (string, []float64, error, error) {
	var choiceOperations string
	fmt.Println(`Укажите какую операцию вы хотите совершить:
AVG - Среднее число транзакций 
SUM - Сумму транзакций
MED - медиану транзакций`)
	fmt.Scan(&choiceOperations)
	if choiceOperations != "AVG" && choiceOperations != "SUM" && choiceOperations != "MED" {
		return "", nil, nil, errors.New("ERROR1:")
	}
	fmt.Println("Введите через запятую ваши транзакции (!Важно, обязательно после каждой транзакции ставить запятую, а для последней не нужно)")
	slice := []float64{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	transaction := scanner.Text()
	operation := strings.Split(transaction, ",")
	if len(operation) == 1 {
		return "", nil, nil, errors.New("ERROR2:")
	}
	for index, value := range operation {
		operation[index] = strings.TrimSpace(value)
		numbers, _ := strconv.ParseFloat(operation[index], 64)
		slice = append(slice, numbers)
	}
	return choiceOperations, slice, nil, nil
}
func Count(choice string, transactions []float64) float64 {
	lens := float64(len(transactions))
	sum := 0.0
	switch choice {
	case "AVG":
		for _, value := range transactions {
			sum += value
		}
		return sum / lens
	case "SUM":
		for _, value := range transactions {
			sum += value
		}
		return sum
	case "MED":
		slices.Sort(transactions)
		halfLen := len(transactions) / 2
		if len(transactions) == 2 {
			return (transactions[0] + transactions[1]) / 2
		} else if len(transactions)%2 != 0 {
			return transactions[halfLen]
		} else {
			return (transactions[halfLen] + transactions[halfLen-1]) / 2
		}
	}
	return 0
}
func choiceUsers() bool {
	var choice string
	fmt.Println("Желаете ли вы повторить расчет, укажите yes или no")
	fmt.Scan(&choice)
	if choice == "yes" || choice == "YES" || choice == "y" || choice == "Y" {
		return true
	}
	return false
}
