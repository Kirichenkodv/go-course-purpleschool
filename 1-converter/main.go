package main

import (
    "errors"
    "fmt"
    "strings"
)

const (
    usd = "usd"
    eur = "eur"
    rub = "rub"
)

type Rates map[string]float64

// RatesToUSD хранит курс каждой валюты к базовой валюте (USD).
// Пример: 1 EUR = 1.176470588 USD (если 1 USD = 0.85 EUR)
var RatesToUSD = &Rates{
    usd: 1.0,       // базовая
    eur: 1 / 0.85,  // 1 USD = 0.85 EUR => 1 EUR = 1/0.85 USD
    rub: 1 / 80.44, // 1 USD = 80.44 RUB => 1 RUB = 1/80.44 USD
}

// Нормализация строки валюты
func norm(s string) string {
    return strings.ToLower(strings.TrimSpace(s))
}

// Список доступных валют (для подсказок в UI)
func availableCurrencies(rates *Rates) []string {
    m := *rates
    out := make([]string, 0, len(m))
    for k := range m {
        out = append(out, k)
    }
    return out
}

func printCurrencies(rates *Rates) {
    fmt.Println(strings.Join(availableCurrencies(rates), " "))
}

// Конвертация через базовую валюту (USD):
// amount[src] -> USD -> amount[tgt]
func convert(amount float64, src, tgt string, rates *Rates) (float64, error) {
    src = norm(src)
    tgt = norm(tgt)

    m := *rates
    rateSrcUSD, okSrc := m[src]
    rateTgtUSD, okTgt := m[tgt]
    if !okSrc || !okTgt {
        return 0, errors.New("неизвестная валюта")
    }
    if amount <= 0 {
        return 0, errors.New("сумма должна быть больше 0")
    }
    // amount[src] * (USD per src) / (USD per tgt) = amount[tgt]
    return amount * rateSrcUSD / rateTgtUSD, nil
}

func inputAmount(rates *Rates) (sourceCurrency string, amount float64, targetCurrency string) {
    m := *rates

    // 1) Ввод исходной валюты
    for {
        fmt.Println("Введите наименование начальной валюты из следующих возможных:")
        printCurrencies(rates)
        if _, err := fmt.Scan(&sourceCurrency); err != nil {
            fmt.Println("Ошибка ввода. Повторите.")
            continue
        }
        sourceCurrency = norm(sourceCurrency)
        if _, ok := m[sourceCurrency]; ok {
            break
        }
        fmt.Println("Неизвестная валюта. Повторите ввод.")
    }
    fmt.Println("Вы выбрали валюту", sourceCurrency)

    // 2) Ввод суммы (> 0)
    for {
        fmt.Println("Введите сумму (> 0):")
        if _, err := fmt.Scan(&amount); err != nil {
            fmt.Println("Ошибка чтения суммы. Повторите.")
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
        printCurrencies(rates)
        if _, err := fmt.Scan(&targetCurrency); err != nil {
            fmt.Println("Ошибка ввода. Повторите.")
            continue
        }
        targetCurrency = norm(targetCurrency)
        if targetCurrency == sourceCurrency {
            fmt.Println("Целевая валюта совпадает с исходной. Повторите.")
            continue
        }
        if _, ok := m[targetCurrency]; ok {
            break
        }
        fmt.Println("Недопустимая целевая валюта. Повторите.")
    }

    fmt.Println("Вам надо переконвертировать", amount, sourceCurrency)
    fmt.Println("Валюта для конвертации", targetCurrency)
    return
}

func main() {
    src, amt, tgt := inputAmount(RatesToUSD)

    res, err := convert(amt, src, tgt, RatesToUSD)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Println("Результат:", res)
}
