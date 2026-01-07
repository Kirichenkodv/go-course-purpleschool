package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func readLineCSV() ([]float64, error) {
	fmt.Print("Введите числа через запятую и нажмите Enter: ")
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return nil, err
	}
	return splitNumbersByComma(line), nil
}

func splitNumbersByComma(s string) []float64 {
	parts := strings.Split(s, ",")
	nums := make([]float64, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			if num, err := strconv.ParseFloat(p, 64); err == nil {
				nums = append(nums, num)
			}
		}
	}
	return nums
}

func main() {
	operation := selectOperation()
	fmt.Println("Вы выбрали:", operation)

	numbers, err := readLineCSV()
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Println("Введено:", numbers)
}
