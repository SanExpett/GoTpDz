package main

import (
	"fmt"

	"github.com/SanExpett/TpGoDz/calculate"
)

func main() {
	err := calculate.Run()
	if err != nil {
		fmt.Printf("Произошла ошибка: %v", err)
	}
}
