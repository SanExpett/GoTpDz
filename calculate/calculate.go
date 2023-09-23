package calculate

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"fmt"
)

func Run() error {
	expression, err := parceCommandLine()
	if err != nil {
		return err
	}

	result, err := calculate(expression)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

func parceCommandLine() (string, error){ // вытаскиваем из строки выражение
	if len(os.Args) > 2 {
		return "", errors.New("There must be only one expression")
	}

	if len(os.Args) == 1 {
		return "", errors.New("No expression")
	}

	expression := os.Args[1]
	return expression, nil
}

func calculate(expression string) (string, error) { // передаем строчку с выражением, получаем итоговый ответ строкой
	tokens, err := tokenize(expression)
	if err != nil {
		return "", err
	}

	inPol, err := ToPolishNotation(tokens)
	if err != nil {
		return "", err
	}

	result, err := getResFromPolish(inPol)
	if err != nil {
		return "", err
	}

	return result, nil
}

func calcForTwoNums(num1 string, num2 string, operator string) (string, error) { // передаем стрки с двумя числами и операцией, получаем резульат ее применения на числах
	intNum1, _ := strconv.Atoi(num1)
	intNum2, _ := strconv.Atoi(num2)

	switch {
	case operator == "+":
		return strconv.Itoa(intNum1 + intNum2), nil
	case operator == "-":
		return strconv.Itoa(intNum1 - intNum2), nil
	case operator == "*":
		return strconv.Itoa(intNum1 * intNum2), nil
	case operator == "/":
		return strconv.Itoa(intNum1 / intNum2), nil
	default:
		return "", nil
	}
}

func tokenize(expression string) ([]string, error) { // из строки с выражением в слайс с числами и операторами
	if !correctExpression(expression) {
		return nil, errors.New("Only digits and signs")
	}

	tokens := []string{}
	token := ""

	for _, char := range expression {
		switch {
		case char >= '0' && char <= '9':
			token += string(char)
		case char == '+' || char == '-' || char == '*' || char == '/' || char == '(' || char == ')':
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
			tokens = append(tokens, string(char))
		case char == ' ':
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
		default:
			return []string{}, nil
		}
	}

	if token != "" {
		tokens = append(tokens, token)
	}

	return tokens, nil
}

func correctExpression(expr string) bool { // проверяем что выражение состоит из цифр, скобок и пробелов
	reg, _ := regexp.Compile(`^[0-9()+\-*\/ ]+$`)
	return reg.MatchString(expr)
}
