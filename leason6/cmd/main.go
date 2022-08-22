package main

import (
	"GBGo/leason6/node"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var listNode = new(node.ListNode)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Не число. Пропускаем")
			continue
		}

		listNode.Add(input)
	}

	fmt.Println(listNode.GetAll())
	fmt.Println(listNode.GetAllv2())
}
