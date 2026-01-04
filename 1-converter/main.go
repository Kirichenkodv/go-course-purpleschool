package main

import "fmt"

func main() {
	const (
		USDToEUR = 0.85
		USDToRUB = 80.44
	)
	eurToRub := (1.0 / USDToEUR) * USDToRUB
	fmt.Println("Курс евро к рублю составляет")
	fmt.Println(eurToRub)
}
