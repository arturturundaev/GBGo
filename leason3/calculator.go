package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float32
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число (для рассчета факториала - необязательное): ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, f): ")
	fmt.Scanln(&op)
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Нельзя делить на 0")
			os.Exit(0)
		}
		res = a / b
	case "^":
		res = float32(math.Pow(float64(a), float64(b)))

	case "f":
		res = float32(calculateFactorial(int(a)))

	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func calculateFactorial(n int) uint64 {
	var factVal uint64 = 1

	if n < 0 {
		fmt.Print("Нельзя рассчитать факториал от отрицательного числа")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
	}
	return factVal
}
