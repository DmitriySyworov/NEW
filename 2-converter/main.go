package main

import "fmt"

func main() {
	number, orig, targ := output()
	rez := countCurrencies(number, orig, targ)

}
func output() (float64, string, string) {
	var numberCurrency float64
	var original, target string
	fmt.Scan(&numberCurrency, &original, &target)
	return numberCurrency, original, target
}
func countCurrencies(numbers float64, Original, Target string) float64 {

}
