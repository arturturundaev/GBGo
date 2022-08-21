package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var fibonacciMap = map[int]int{
	0: 0,
	1: 1,
}

// Количество раз, сколько напрямую пришлось рассчитывать значение числа, а не брать от мапы
var countCalculate int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Не число. Пропускаем")
			continue
		}
		result := fibonacci(input)
		fmt.Println("Сумма: ", result, " Количество прямых расчетов (не через кеш) :", countCalculate)
		countCalculate = 0
	}
}

/**
 * Рассчет суммы числа Фибоначчи
 * https://ru.wikipedia.org/wiki/%D0%A7%D0%B8%D1%81%D0%BB%D0%B0_%D0%A4%D0%B8%D0%B1%D0%BE%D0%BD%D0%B0%D1%87%D1%87%D0%B8
 */
func fibonacci(number int) int {
	if number <= 1 {
		return number
	}

	if _, ok := fibonacciMap[number]; ok {
		return fibonacciMap[number]
	}

	countCalculate++
	fibonacciMap[number] = fibonacci(number-1) + fibonacci(number-2) + 1

	return fibonacciMap[number]
}
