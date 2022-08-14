package main

import (
	"fmt"
	"os"
)

func main() {
	var data int
	var first, second, third int

	fmt.Println("enter digital")
	_, err := fmt.Scanf("%d\n", &data)
	if err != nil {
		println("Fscanf: ", err.Error())
		os.Exit(0)
	}

	first = data / 100
	data = data % 100

	second = data / 10
	third = data % 10

	println(first, " сотни ", second, " десятка и ", third, " едениц")
}
