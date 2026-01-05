package main

import "fmt"

const (
	USDToEUR = 0.85
	USDToRUB = 80.44
)
const (
	usd = "usd"
	eur = "eur"
	rub = "rub"
)

func inputAmount() (sourceCurrency string, amount int, targetCurrency string) {
	// 1) Ввод исходной валюты
	for {
		fmt.Println("Введите наименование начальной валюты из следующих возможных:")
		fmt.Println(usd, eur, rub)
		if _, err := fmt.Scan(&sourceCurrency); err != nil {
			fmt.Println("Ошибка ввода. Повторите.")
			continue
		}
		if sourceCurrency == usd || sourceCurrency == eur || sourceCurrency == rub {
			break
		}
		fmt.Println("Неизвестная валюта. Повторите ввод.")
	}
	fmt.Println("Вы выбрали валюту", sourceCurrency)

	// 2) Ввод суммы (целое число > 0)
	for {
		fmt.Println("Введите сумму (целое число > 0):")
		if _, err := fmt.Scan(&amount); err != nil {
			fmt.Println("Ошибка чтения суммы. Повторите.")
			// сбросим amount на случай частичного ввода
			amount = 0
			continue
		}
		if amount > 0 {
			break
		}
		fmt.Println("Сумма должна быть больше 0. Повторите.")
	}

	// 3) Ввод целевой валюты (исключая исходную)
	for {
		fmt.Println("Введите наименование конечной валюты из следующих возможных:")
		switch sourceCurrency {
		case usd:
			fmt.Println(eur, rub)
		case eur:
			fmt.Println(usd, rub)
		case rub:
			fmt.Println(usd, eur)
		}

		if _, err := fmt.Scan(&targetCurrency); err != nil {
			fmt.Println("Ошибка ввода. Повторите.")
			continue
		}

		// проверка допустимости и исключение совпадения
		valid := false
		switch sourceCurrency {
		case usd:
			valid = (targetCurrency == eur || targetCurrency == rub)
		case eur:
			valid = (targetCurrency == usd || targetCurrency == rub)
		case rub:
			valid = (targetCurrency == usd || targetCurrency == eur)
		}

		if !valid {
			fmt.Println("Недопустимая целевая валюта для выбранной исходной. Повторите.")
			continue
		}
		break
	}

	fmt.Println("Вам надо переконвертировать", amount, sourceCurrency)
	fmt.Println("Валюта для конвертации", targetCurrency)
	return
}

func convertAmount(
	sourceCurrency string,
	amount int,
	targetCurrency string) {
	switch {
	case sourceCurrency == usd:
		switch targetCurrency {
		case eur:
			fmt.Println("Результат:", float64(amount)*USDToEUR)
		case rub:
			fmt.Println("Результат:", float64(amount)*USDToRUB)
		}
	case sourceCurrency == eur:
		switch targetCurrency {
		case usd:
			fmt.Println("Результат:", float64(amount)/USDToEUR)
		case rub:
			fmt.Println("Результат:", float64(amount)*USDToRUB/USDToEUR)
		}
	case sourceCurrency == rub:
		switch targetCurrency {
		case usd:
			fmt.Println("Результат:", float64(amount)/USDToRUB)
		case eur:
			fmt.Println("Результат:", float64(amount)/USDToRUB*USDToEUR)
		}
	}
}

func main() {

	sourceCurrency, amount, targetCurrency := inputAmount()
	convertAmount(sourceCurrency, amount, targetCurrency)
}
