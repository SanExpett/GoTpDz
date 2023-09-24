package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPolishNotation(t *testing.T) {
	in := []string{"(", "1", "+", "23", ")", "*", "33"}
	expected := []string{"1", "23", "+", "33", "*"}
	result, _ := ToPolishNotation(in)
	assert.Equal(t, expected, result)

	in = []string{"(", "14", "+", "16", "-", "3", ")", "/", "(", "3", "*", "3", ")"}
	expected = []string{"14", "16", "+", "3", "-", "3", "3", "*", "/"}
	result, _ = ToPolishNotation(in)
	assert.Equal(t, expected, result)
}

func TestToPolishNotationErrors(t *testing.T) {
	in := []string{")", "1", "+", "23", ")", "*", "33"}
	_, err := ToPolishNotation(in)
	assert.NotNil(t, err)

	in = []string{"(", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "+", "23", ")", "*", "33"}
	_, err = ToPolishNotation(in)
	assert.NotNil(t, err)
}

func TestCulculate(t *testing.T) {
	in := "(1+2)-3"
	actual, _ := calculate(in)
	assert.Equal(t, "0", actual)

	in = "(1+2)*3"
	actual, _ = calculate(in)
	assert.Equal(t, "9", actual)

	in = "(14+16-3)/(3*3)"
	actual, _ = calculate(in)
	assert.Equal(t, "3", actual)
}

func TestGetResFromPolish(t *testing.T) {
	in := []string{"14", "16", "+", "3", "-", "3", "3", "*", "/"}
	actual, _ := getResFromPolish(in)
	expected := "3"
	assert.Equal(t, expected, actual)
}

func TestTokenize(t *testing.T) {
	in := "(14+16-3)/(3*3)"
	actual, _ := tokenize(in)
	expected := []string{"(", "14", "+", "16", "-", "3", ")", "/", "(", "3", "*", "3", ")"}
	assert.Equal(t, expected, actual)
}

func TestCalcForTwoNums(t *testing.T) {
	num1 := "25"
	num2 := "5"

	op := "+"
	actual, _ := calcForTwoNums(num1, num2, op)
	expected := "30"
	assert.Equal(t, expected, actual)

	op = "-"
	actual, _ = calcForTwoNums(num1, num2, op)
	expected = "20"
	assert.Equal(t, expected, actual)

	op = "*"
	actual, _ = calcForTwoNums(num1, num2, op)
	expected = "125"
	assert.Equal(t, expected, actual)

	op = "/"
	actual, _ = calcForTwoNums(num1, num2, op)
	expected = "5"
	assert.Equal(t, expected, actual)
}

func TestCorrectExpression(t *testing.T) {
	in := "(14+16-3)/(3*3)"
	expected := true
	actual := correctExpression(in)
	assert.Equal(t, expected, actual)

	in = "(14+16-3)/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA(3*3)"
	expected = false
	actual = correctExpression(in)
	assert.Equal(t, expected, actual)
}

func TestFixExpression(t *testing.T) {
	in := "-1+2"
	expected := "0-1+2"
	actual := fixExpression(in)
	assert.Equal(t, expected, actual)

	in = "2+(-3+4)"
	expected = "2+(0-3+4)"
	actual = fixExpression(in)
	assert.Equal(t, expected, actual)
}
