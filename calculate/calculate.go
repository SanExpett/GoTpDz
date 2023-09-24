package calculate

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() error {
	expression, err := parceCommandLine()
	if err != nil {
		return err
	}

	expression = fixExpression(expression)

	result, err := calculate(expression)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

// учитывает ситуацию с отрицательными числами, если перед символом минус ничего нет(начало строки)
// или (, то это минус от отрицательного числа, функция заменяет такие - на 0-.
func fixExpression(str string) string {
	var result []string

	for i := 0; i < len(str); i++ {
		if str[i] == '-' {
			if i == 0 || str[i-1] == '(' {
				result = append(result, "0-")
			} else {
				result = append(result, string(str[i]))
			}
		} else {
			result = append(result, string(str[i]))
		}
	}

	return strings.Join(result, "")
}

func parceCommandLine() (string, error) { // вытаскиваем из строки выражение
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

// передаем стрки с двумя числами и операцией, получаем резульат ее применения на числах.
func calcForTwoNums(num1 string, num2 string, operator string) (string, error) { 
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
		if num2 == "0" {
			return "0", errors.New("Zero division")
		}
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
	reg := regexp.MustCompile(`^[0-9()+\-*\/ ]+$`)
	return reg.MatchString(expr)
}
