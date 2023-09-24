package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPolishNotation(t *testing.T) {
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
	input := []string{")", "1", "+", "23", ")", "*", "33"}
	_, err := ToPolishNotation(input)
	assert.NotNil(t, err)

	input = []string{"(", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "+", "23", ")", "*", "33"}
	_, err = ToPolishNotation(input)
	assert.NotNil(t, err)
}

func TestCulculate(t *testing.T) {
	input := "(1+2)-3"
	actual, _ := calculate(input)
	assert.Equal(t, "0", actual)

	input = "(1+2)*3"
	actual, _ = calculate(input)
	assert.Equal(t, "9", actual)

	input = "(14+16-3)/(3*3)"
	actual, _ = calculate(input)
	assert.Equal(t, "3", actual)
}

func TestGetResFromPolish(t *testing.T) {
	input := []string{"14", "16", "+", "3", "-", "3", "3", "*", "/"}
	actual, _ := getResFromPolish(input)
	expected := "3"
	assert.Equal(t, expected, actual)
}

func TestTokenize(t *testing.T) {
	input := "(14+16-3)/(3*3)"
	actual, _ := tokenize(input)
	expected := []string{"(", "14", "+", "16", "-", "3", ")", "/", "(", "3", "*", "3", ")"}
	assert.Equal(t, expected, actual)
}

func TestCalcForTwoNums(t *testing.T) {
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
	input := "(14+16-3)/(3*3)"
	expected := true
	actual := correctExpression(input)
	assert.Equal(t, expected, actual)

	input = "(14+16-3)/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA(3*3)"
	expected = false
	actual = correctExpression(input)
	assert.Equal(t, expected, actual)
}

func TestFixExpression(t *testing.T) {
	input := "-1+2"
	expected := "0-1+2"
	actual := fixExpression(input)
	assert.Equal(t, expected, actual)

	input = "2+(-3+4)"
	expected = "2+(0-3+4)"
	actual = fixExpression(input)
	assert.Equal(t, expected, actual)
}
