package main

import "fmt"

func inputAmount() {
	var (
		amount         int
		sourceCurrency string
		targetCurrency string
	)
	fmt.Scanln(&amount, &sourceCurrency, &targetCurrency)
	fmt.Println(amount, sourceCurrency, targetCurrency)
}

func convertAmount(amount int, sourceCurrency, targetCurrency string) {

}

func main() {
	const (
		USDToEUR = 0.85
		USDToRUB = 80.44
	)
	eurToRub := (1.0 / USDToEUR) * USDToRUB
	fmt.Println("Курс евро к рублю составляет")
	fmt.Println(eurToRub)
	inputAmount()
}
