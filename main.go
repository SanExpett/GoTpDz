package main

import (
	"fmt"
	"os"
	"errors"

	"github.com/SanExpett/TpGoDz/calculate"
)

var errParce = errors.New("error in parcing command line")

func parceCommandLine() (string, error) { // вытаскиваем из строки выражение
	if len(os.Args) > 2 {
		return "", errParce
	}

	if len(os.Args) == 1 {
		return "", errParce
	}

	expression := os.Args[1]

	return expression, nil
}

func main() {
	expression, err := parceCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err := calculate.Calculate(expression)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
