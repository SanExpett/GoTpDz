package functions_test

import (
	"testing"

	"github.com/SanExpett/GoDz1P1/functions"
	"github.com/stretchr/testify/assert"
)

var EXAMPLE = []string{
	"I love music.",
	"I love music.",
	"I love music.",
	"",
	"I love music of Kartik.",
	"I love music of Kartik.",
	"Thanks.",
	"I love music of Kartik.",
	"I love music of Kartik.",
}

func TestUniqLines(t *testing.T) {
	t.Parallel()

	lines := EXAMPLE
	result := functions.UniqLines(lines)
	expected := []bool{true, false, false, true, true, false, true, true, false}
	assert.Equal(t, expected, result)
}

func TestCountOfLines(t *testing.T) {
	t.Parallel()

	// тестим то, что функция правильно возвращает слайс булов с подходящими строками
	lines := EXAMPLE // предполагается, что это строки которые возможно были поломаны ignore флагами

	linesWithCounts := make([]string, len(lines))
	copy(linesWithCounts, lines)

	result := functions.CountOfLines(lines, &linesWithCounts)
	expectedBools := []bool{true, false, false, true, true, false, true, true, false}
	assert.Equal(t, result, expectedBools)

	// тестим, что функция правильно считает количество повторяющихся строк
	expectedLines := []string{
		"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik.",
	}

	var resultLines []string

	for i, val := range expectedBools {
		if val {
			resultLines = append(resultLines, linesWithCounts[i])
		}
	}

	assert.Equal(t, expectedLines, resultLines)
}

func TestRepeatingLines(t *testing.T) {
	t.Parallel()

	lines := EXAMPLE
	result := functions.RepeatingLines(lines)
	expected := []bool{true, false, false, false, true, false, false, true, false}
	assert.Equal(t, expected, result)
}

func TestNonRepeatingLines(t *testing.T) {
	t.Parallel()

	lines := EXAMPLE
	result := functions.NonRepeatingLines(lines)
	expected := []bool{false, false, false, true, false, false, true, false, false}
	assert.Equal(t, expected, result)
}

func TestGetResultRignt(t *testing.T) {
	t.Parallel()

	lines := EXAMPLE
	suitableLines := []bool{false, false, false, true, false, false, true, false, false}
	_ = functions.GetResult(&lines, suitableLines)

	expected := []string{"", "Thanks."}
	assert.Equal(t, expected, lines)
}

func TestGetResultError(t *testing.T) {
	t.Parallel()

	lines := EXAMPLE
	suitableLines := []bool{false, false, false}
	err := functions.GetResult(&lines, suitableLines)
	assert.NotNil(t, err)
}
