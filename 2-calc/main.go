package main

import "fmt"

const (
	AVG = "avg"
	SUM = "sum"
	MED = "med"
)

func selectOperation() (operation string) {
	for {
		fmt.Println("Введите название операции из следующих возможных")
		fmt.Println(AVG, SUM, MED)
		if _, err := fmt.Scan(&operation); err != nil {
			fmt.Println("Ошибка ввода. Повторите.")
			continue
		}
		if operation == AVG || operation == SUM || operation == MED {
			break
		}
		fmt.Println("Нет такой операции")

	}
	return
}

func main() {
	operation := selectOperation()
	fmt.Println("Вы выбрали:", operation)
}
