package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputData := make([]int, 0)
	scaner := bufio.NewScanner(os.Stdin)

	for scaner.Scan() {
		input, err := strconv.Atoi(scaner.Text())
		if err != nil {
			fmt.Println("Не число. Пропускаем")
			continue
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
