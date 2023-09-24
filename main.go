package main

import (
	"fmt"

	"github.com/SanExpett/TpGoDz/calculate"
)

func main() {
	err := calculate.Run()
	if err != nil {
		errorMsg := fmt.Sprintf("Произошла ошибка: %v", err)
		fmt.Println(errorMsg)
	}
}
