package functions

import (
	"testing"

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
	lines := EXAMPLE
	result := UniqLines(lines)
	expected := []bool{true, false, false, true, true, false, true, true, false}
	assert.Equal(t, result, expected)
}

func TestCountOfLines(t *testing.T) {
	// тестим то, что функция правильно возвращает слайс булов с подходящими строками
	lines := EXAMPLE // предполагается, что это строки которые возможно были поломаны ignore флагами
	linesWithCounts := EXAMPLE
	result := CountOfLines(lines, &linesWithCounts)
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
	assert.Equal(t, resultLines, expectedLines)
}

func TestRepeatingLines(t *testing.T) {
	lines := EXAMPLE
	result := RepeatingLines(lines)
	expected := []bool{true, false, false, false, true, false, false, true, false}
	assert.Equal(t, result, expected)
}

func TestNonRepeatingLines(t *testing.T) {
	lines := EXAMPLE
	result := NonRepeatingLines(lines)
	expected := []bool{false, false, false, true, false, false, true, false, false}
	assert.Equal(t, result, expected)
}

func TestGetResult(t *testing.T) {
	lines := EXAMPLE
	result := NonRepeatingLines(lines)
	expected := []bool{false, false, false, true, false, false, true, false, false}
	assert.Equal(t, result, expected)
}

