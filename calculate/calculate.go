package calculate

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)


var errCalc  = errors.New("calculating error")

// передаем строчку с выражением, получаем итоговый ответ строкой.
func Calculate(expression string) (string, error) {
	tokens, err := tokenize(expression)
	if err != nil {
		return "", err
	}

	inPol, err := toPolishNotation(tokens)
	if err != nil {
		return "", err
	}

	result, err := getResFromPolish(inPol)
	if err != nil {
		return "", err
	}

	return result, nil
}

// передаем стрки с двумя числами и операцией, получаем резульат ее применения на числах.
func calcForTwoNums(num1 string, num2 string, operator string) (string, error) {
	floatNum1, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return "", fmt.Errorf(errTemplate, err)
	}

	floatNum2, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return "", fmt.Errorf(errTemplate, err)
	}

	switch {
	case operator == "+":
		return strconv.FormatFloat(floatNum1+floatNum2, 'f', -1, 64), nil
	case operator == "-":
		return strconv.FormatFloat(floatNum1-floatNum2, 'f', -1, 64), nil
	case operator == "*":
		return strconv.FormatFloat(floatNum1*floatNum2, 'f', -1, 64), nil
	case operator == "/":
		if floatNum2 == 0.0 {
			return "0", errCalc
		}

		return strconv.FormatFloat(floatNum1/floatNum2, 'f', -1, 64), nil
	default:
		return "", nil
	}
}

func tokenize(expression string) ([]string, error) { // из строки с выражением в слайс с числами и операторами
	if !correctExpression(expression) {
		return nil, errCalc
	}

	tokens := []string{}
	token := ""

	for _, char := range expression {
		switch {
		case char >= '0' && char <= '9' || char == '.':
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
	reg := regexp.MustCompile(`^[0-9()+\-*\/. ]+$`)

	return reg.MatchString(expr)
}
