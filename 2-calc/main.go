package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	AVG = "avg"
	SUM = "sum"
	MED = "med"
)

func main() {
	operation := selectOperation()
	numbers, err := readLineCSV()
	if err != nil {
		printError(err)
		return
	}
	if len(numbers) == 0 {
		printError(errors.New("не введено ни одного корректного числа"))
		return
	}
	printSummary(operation, numbers)
	printResult(operation, numbers)
}

/* ---------- Ввод ---------- */

func selectOperation() (operation string) {
	for {
		printHeader("Выбор операции")
		fmt.Println("Доступные операции:")
		fmt.Printf("  • %s — среднее арифметическое\n", AVG)
		fmt.Printf("  • %s — сумма\n", SUM)
		fmt.Printf("  • %s — медиана\n", MED)
		fmt.Print("\nВведите название операции: ")

		if _, err := fmt.Scan(&operation); err != nil {
			fmt.Println("Ошибка ввода. Повторите.")
			continue
		}
		operation = strings.ToLower(strings.TrimSpace(operation))
		if operation == AVG || operation == SUM || operation == MED {
			fmt.Println("Вы выбрали:", operation)
			return
		}
		fmt.Println("Нет такой операции. Повторите.\n")
	}
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
		if p == "" {
			continue
		}
		if num, err := strconv.ParseFloat(p, 64); err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

/* ---------- Вычисления ---------- */

func calcSum(numbers []float64) (sum float64) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func calcAvg(numbers []float64) float64 {
	return calcSum(numbers) / float64(len(numbers))
}

func calcMed(numbers []float64) float64 {
	// Не меняем исходный слайс
	cp := make([]float64, len(numbers))
	copy(cp, numbers)
	sort.Float64s(cp)

	n := len(cp)
	if n%2 == 1 {
		return cp[n/2]
	}
	return (cp[n/2-1] + cp[n/2]) / 2.0
}

/* ---------- Красивый вывод ---------- */

func printHeader(title string) {
	sep := strings.Repeat("─", 40)
	fmt.Printf("\n%s\n%s\n", title, sep)
}

func printError(err error) {
	printHeader("Ошибка")
	fmt.Println("  ", err)
}

func formatNumber(v float64) string {
	// Компактный и читабельный формат
	return fmt.Sprintf("%.6g", v)
}

func formatSlice(nums []float64) string {
	if len(nums) == 0 {
		return "[]"
	}
	var b strings.Builder
	b.WriteString("[")
	for i, v := range nums {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(formatNumber(v))
	}
	b.WriteString("]")
	return b.String()
}

func printSummary(operation string, numbers []float64) {
	printHeader("Введённые данные")
	fmt.Println("  Операция:", operation)
	fmt.Println("  Кол-во чисел:", len(numbers))
	fmt.Println("  Числа:", formatSlice(numbers))
}

func printResult(operation string, numbers []float64) {
	printHeader("Результат")
	switch operation {
	case AVG:
		fmt.Println("  Среднее:", formatNumber(calcAvg(numbers)))
	case SUM:
		fmt.Println("  Сумма:  ", formatNumber(calcSum(numbers)))
	case MED:
		fmt.Println("  Медиана:", formatNumber(calcMed(numbers)))
	}
}
