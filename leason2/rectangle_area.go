package main

import (
	"fmt"
	"os"
)

func main() {
	var length int
	var width int

	fmt.Println("enter length")
	_, err := fmt.Scanf("%d\n", &length)
	if err != nil {
		println("Fscanf: ", err.Error())
		os.Exit(0)
	}

	fmt.Println("enter width")

	_, err = fmt.Scanf("%d\n", &width)
	if err != nil {
		println("Fscanf: ", err.Error())
		os.Exit(0)
	}

	result := calculateRectangleArea(length, width)
	println("Result: ", result)
}

func calculateRectangleArea(length int, width int) int {

	return length * width
}
