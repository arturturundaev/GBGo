package main

import (
	"config/configuration"
	"fmt"
	"os"
)

func main() {
	config, err := fmt.Println(configuration.GetConfiguration())
	if err != nil {
		fmt.Println("can't process the config: %w", err)
		os.Exit(0)
	}

	fmt.Println(config)
}
