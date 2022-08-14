package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Print("Введите число: ")
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		println("Fscanf: ", err.Error())
		os.Exit(0)
	}

	result := simpleNumbers(n)
	fmt.Println(result)
}

func simpleNumbers(n int) []int {
	var data []int // Массив простых чисел
	var canBySimple bool = false

	if n <= 1 {
		fmt.Println("Число долно быть больше 1")
		return data
	}

	data = append(data, 2)
	for i := 3; i <= n; i++ {
		canBySimple = true
		for _, value := range data {
			// Число из первого цила не должно делиться нацело
			// Исключаем деление на 1 и на самого себя
			if i%value == 0 && value != 1 && value != i {
				canBySimple = false
			}
		}

		if canBySimple {
			data = append(data, i)
		}
	}

	return data
}
