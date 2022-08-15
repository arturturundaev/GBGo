package main

import (
	"fmt"
)

func main() {
	var input int
	inputData := make([]int, 0)
	fmt.Println("Для завершения ввода введите любой символ\n")
	for {
		fmt.Print("Введите значение: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			break
		}
		inputData = insertWithSort(inputData, input)
	}

	fmt.Println(inputData)
}

func insertWithSort(inputData []int, input int) []int {
	inputData = append(inputData, input)

	if len(inputData) == 1 {
		return inputData
	}

	var tmp int
	for i := 1; i < len(inputData); i++ {
		for j := i; j > 0 && inputData[j-1] > inputData[j]; j-- {
			tmp = inputData[j-1]
			inputData[j-1] = inputData[j]
			inputData[j] = tmp
		}
	}

	return inputData
}
