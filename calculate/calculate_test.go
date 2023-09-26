package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPolishNotation(t *testing.T) {
	t.Parallel()

	input := []string{"(", "1", "+", "23", ")", "*", "33"}
	expected := []string{"1", "23", "+", "33", "*"}
	result, _ := ToPolishNotation(input)
	assert.Equal(t, expected, result)

	input = []string{"(", "14", "+", "16", "-", "3", ")", "/", "(", "3", "*", "3", ")"}
	expected = []string{"14", "16", "+", "3", "-", "3", "3", "*", "/"}
	result, _ = ToPolishNotation(input)
	assert.Equal(t, expected, result)
}

func TestToPolishNotationErrors(t *testing.T) {
	t.Parallel()

	input := []string{")", "1", "+", "23", ")", "*", "33"}
	_, err := ToPolishNotation(input)
	assert.NotNil(t, err)

	input = []string{"(", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "+", "23", ")", "*", "33"}
	_, err = ToPolishNotation(input)
	assert.NotNil(t, err)
}

func TestCalculate(t *testing.T) {
	t.Parallel()

	input := "(1+2)-3"
	actual, _ := calculate(input)
	assert.Equal(t, "0", actual)

	input = "(1+2)*3"
	actual, _ = calculate(input)
	assert.Equal(t, "9", actual)

	input = "(14+16-3)/(3*3)"
	actual, _ = calculate(input)
	assert.Equal(t, "3", actual)

	input = "10*-1"
	actual, _ = calculate(input)
	assert.Equal(t, "-10", actual)

	input = "-3+10"
	actual, _ = calculate(input)
	assert.Equal(t, "7", actual)

	input = "-(-4*-10)"
	actual, _ = calculate(input)
	assert.Equal(t, "-40", actual)
}

func TestGetResFromPolish(t *testing.T) {
	t.Parallel()

	input := []string{"14", "16", "+", "3", "-", "3", "3", "*", "/"}
	actual, _ := getResFromPolish(input)
	expected := "3"
	assert.Equal(t, expected, actual)
}

func TestTokenize(t *testing.T) {
	t.Parallel()

	input := "(14+16-3)/(3*3)"
	actual, _ := tokenize(input)
	expected := []string{"(", "14", "+", "16", "-", "3", ")", "/", "(", "3", "*", "3", ")"}
	assert.Equal(t, expected, actual)
}

func TestCalcForTwoNums(t *testing.T) {
	t.Parallel()

	num1 := "25"
	num2 := "5"

	operator := "+"
	actual, _ := calcForTwoNums(num1, num2, operator)
	expected := "30"
	assert.Equal(t, expected, actual)

	operator = "-"
	actual, _ = calcForTwoNums(num1, num2, operator)
	expected = "20"
	assert.Equal(t, expected, actual)

	operator = "*"
	actual, _ = calcForTwoNums(num1, num2, operator)
	expected = "125"
	assert.Equal(t, expected, actual)

	operator = "/"
	actual, _ = calcForTwoNums(num1, num2, operator)
	expected = "5"
	assert.Equal(t, expected, actual)
}

func TestCorrectExpression(t *testing.T) {
	t.Parallel()

	input := "(14+16-3)/(3*3)"
	expected := true
	actual := correctExpression(input)
	assert.Equal(t, expected, actual)

	input = "(14+16-3)/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA(3*3)"
	expected = false
	actual = correctExpression(input)
	assert.Equal(t, expected, actual)
}
