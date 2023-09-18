package main

import (
	"fmt"

	"github.com/SanExpett/TpGoDz/processing"
)

func main() {
	err := processing.ParseCommandLine()
	if err != nil {
		fmt.Printf("Произошла ошибка: %v", err)
	}
}
