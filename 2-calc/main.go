package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func readLineCSV() (string, error) {
	fmt.Print("Введите числа через запятую и нажмите Enter: ")
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && len(line) == 0 {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func main() {
	operation := selectOperation()
	fmt.Println("Вы выбрали:", operation)

	s, err := readLineCSV()
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Println("Введено:", s)
}
