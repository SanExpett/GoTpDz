package calculate

import (
	"errors"
	"github.com/SanExpett/TpGoDz/stack"
	"strconv"
)

func isNum(str string) bool { // проверяем что в строке число
	if _, err := strconv.Atoi(str); err != nil {
		return false
	}
	return true
}

func ToPolishNotation(tokens []string) ([]string, error) { // передаем слайс строк с выражением в стандартной записи, получаем слайс строк в польской нотации
	stack := stack.Create()

	result := make([]string, 0)

	// приоритеты операций
	operators := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for _, token := range tokens {
		switch token {
		case "(":
			stack.Push(token)
		case ")":
			// Пока не встречаем открывающую скобку, переносим операторы из стека в результат
			for top, err := stack.Top(); stack.Len() > 0 && top != "("; top, err = stack.Top() {
				if err != nil {
					return nil, err
				}
				result = append(result, top)
				stack.Pop()
			}

			// Если не удалось найти открывающую скобку, возвращаем ошибку
			top, err := stack.Top()
			if err != nil {
				return nil, err
			}
			if stack.Len() == 0 || top != "(" {
				return nil, errors.New("unbalanced parentheses")
			}

			// Удаляем открывающую скобку из стека
			_, err = stack.Pop()
			if err != nil {
				return nil, err
			}
		default:
			// Если текущий токен - оператор
			if _, isOperator := operators[token]; isOperator {
				// Пока оператор на вершине стека имеет >= приоритет, переносим его в результат
				for top, err := stack.Top(); stack.Len() > 0 && top != "(" && operators[token] <= operators[top]; top, err = stack.Top() {
					if err != nil {
						return nil, err
					}
					result = append(result, top)
					_, err = stack.Pop()
					if err != nil {
						return nil, err
					}
				}

				// Добавляем текущий оператор в стек
				stack.Push(token)
			} else {
				if !isNum(token) {
					return nil, errors.New("Not num")
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
			return nil, err
		}
		if top == "(" {
			return nil, errors.New("unbalanced parentheses")
		}

		result = append(result, top)
		_, err = stack.Pop()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func getResFromPolish(tokens []string) (string, error) { // из выражения в польской записи (слайслом строк) получаем результат выражения строкой
	stack := stack.Create()

	for _, token := range tokens {
		if isNum(token) {
			stack.Push(token)
		}
		if !isNum(token) {
			num2, err := stack.Pop()
			if err != nil {
				return "", err
			}
			num1, err := stack.Pop()
			if err != nil {
				return "", err
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
		return "", err
	}

	return result, nil
}
