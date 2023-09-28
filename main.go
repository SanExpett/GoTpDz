package main

import (
	"fmt"

	"github.com/SanExpett/TpGoDz/calculate"
)

func main() {
	if err := calculate.Run(); err != nil {
		errorMsg := fmt.Errorf("an error has occurred: %w", err)
		fmt.Println(errorMsg)
	}
}
