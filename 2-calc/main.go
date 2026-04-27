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
	operationMap := map[string]func([]float64){
		AVG: calcAvg,
		SUM: calcSum,
		MED: calcMed,
	}

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

	// Вызываем соответствующую функцию расчёта из map
	if operationFunc, exists := operationMap[operation]; exists {
		printHeader("Результат")
		operationFunc(numbers) // Теперь вызываем функцию из map
	} else {
		printError(errors.New("недопустимая операция"))
	}
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
		fmt.Println("Нет такой операции. Повторите.")
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

func calcSum(numbers []float64) {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("  Сумма:  ", formatNumber(sum))
}

func calcAvg(numbers []float64) {
	if len(numbers) == 0 {
		fmt.Println("  Среднее: 0")
		return
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	avg := sum / float64(len(numbers))
	fmt.Println("  Среднее:", formatNumber(avg))
}

func calcMed(numbers []float64) {
	cp := make([]float64, len(numbers))
	copy(cp, numbers)
	sort.Float64s(cp)

	n := len(cp)
	if n%2 == 1 {
		fmt.Println("  Медиана:", formatNumber(cp[n/2]))
	} else {
		median := (cp[n/2-1] + cp[n/2]) / 2.0
		fmt.Println("  Медиана:", formatNumber(median))
	}
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
