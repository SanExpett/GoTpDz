package calculate

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/SanExpett/TpGoDz/stack"
)

var errNotation = errors.New("notation conversion error")

const errTemplate = "%w"

func isInt(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}

	return false
}

func isFloat(str string) bool {
	if _, err := strconv.ParseFloat(str, 32); err == nil {
		return true
	}

	return false
}

func isNum(str string) bool { // проверяем что в строке число
	if isFloat(str) || isInt(str) {
		return true
	}

	return false
}

// nolint: funlen
// nolint: gci
func toPolishNotation(tokens []string) ([]string, error) { // передаем слайс строк с выражением в стандартной записи, получаем слайс строк в польской нотации
	stack := stack.Create()

	result := make([]string, 0)

	// приоритеты операций
	operators := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for idx, token := range tokens {
		switch token {
		case "(":
			stack.Push(token)
		case ")":
			// Пока не встречаем открывающую скобку, переносим операторы из стека в результат
			for top, err := stack.Top(); stack.Len() > 0 && top != "("; top, err = stack.Top() {
				if err != nil {
					return nil, fmt.Errorf(errTemplate, err)
				}

				result = append(result, top)

				_, err = stack.Pop()
				if err != nil {
					return nil, fmt.Errorf(errTemplate, err)
				}
			}

			// Если не удалось найти открывающую скобку, возвращаем ошибку
			top, err := stack.Top()
			if err != nil {
				return nil, fmt.Errorf(errTemplate, err)
			}

			if stack.Len() == 0 || top != "(" {
				return nil, errNotation
			}

			// Удаляем открывающую скобку из стека ?????????????
			_, err = stack.Pop()
			if err != nil {
				return nil, fmt.Errorf(errTemplate, err)
			}
		default:
			// Если текущий токен - оператор
			_, isOperator := operators[token]
			opPrior := operators[token]

			if token == "-" && (idx == 0 || !isNum(tokens[idx-1]) && tokens[idx-1] != ")") { // если унарный минус
				result = append(result, "0")
				opPrior = 3
			}

			if isOperator { // любой другой оператор
				// Пока оператор на вершине стека имеет >= приоритет, переносим его в результат
				for top, err := stack.Top(); stack.Len() > 0 && top != "(" && opPrior <= operators[top]; top, err = stack.Top() {
					if err != nil {
						return nil, fmt.Errorf(errTemplate, err)
					}

					result = append(result, top)

					_, err = stack.Pop()
					if err != nil {
						return nil, fmt.Errorf(errTemplate, err)
					}
				}

				// Добавляем текущий оператор в стек
				stack.Push(token)
			} else {
				if !isNum(token) {
					return nil, errNotation
				}
				// Если текущий токен - число, добавляем его в результат
				result = append(result, token)
			}
		}
	}

	// Если в стеке остались операторы, переносим их в результат
	for stack.Len() > 0 {
		// Если на вершине стека оказывается открывающая скобка => выражение не сбалансировано
		top, err := stack.Top()
		if err != nil {
			return nil, fmt.Errorf(errTemplate, err)
		}

		if top == "(" {
			return nil, errNotation
		}

		result = append(result, top)

		_, err = stack.Pop()
		if err != nil {
			return nil, fmt.Errorf(errTemplate, err)
		}
	}

	return result, nil
}

// из выражения в польской записи (слайслом строк) получаем результат выражения строкой.
func getResFromPolish(tokens []string) (string, error) {
	stack := stack.Create()

	for _, token := range tokens {
		if isNum(token) {
			stack.Push(token)
		}

		if !isNum(token) {
			num2, err := stack.Pop()
			if err != nil {
				return "", fmt.Errorf(errTemplate, err)
			}

			num1, err := stack.Pop()
			if err != nil {
				return "", fmt.Errorf(errTemplate, err)
			}

			res, err := calcForTwoNums(num1, num2, token)
			if err != nil {
				return "", err
			}

			stack.Push(res)
		}
	}

	result, err := stack.Pop()
	if err != nil {
		return "", fmt.Errorf(errTemplate, err)
	}

	return result, nil
}
