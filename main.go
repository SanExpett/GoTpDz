package main

import (
	"fmt"

	"github.com/SanExpett/GoDz1P1/processing"
)

func main() {
	err := processing.ParseCommandLine()
	if err != nil {
		fmt.Printf("Произошла ошибка: %v", err)
	}
}
