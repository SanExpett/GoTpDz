// в этом файле будет лежать код функций для флагов выборки строк (дефолтное поведение, -c, -u, -d)

package functions

import (
	"errors"
	"strconv"
)

// из слайса всех строк и слайса булов подходящих
// вернуть слайс подходящих.
func GetResult(lines *[]string, suitableLines []bool) error {
	if len(*lines) != len(suitableLines) {
		return errors.New("slices must have same len")
	}
	
	var result []string
	for i, val := range suitableLines {
		if val {
			result = append(result, (*lines)[i])
		}
	}
	
	*lines = result

	return nil 
}

// UniqLines возвращает слайс булов, где true обозначает уникальные строки в данном слайсе строк.
func UniqLines(lines []string) []bool { // дефолтное поведение
	if len(lines) == 0 {
		return []bool{}
	}

	suitableLines := make([]bool, len(lines))
	currLine := lines[0]
	suitableLines[0] = true
	for i := 1; i < len(lines); i++ {
		if lines[i] != currLine {
			suitableLines[i] = true
			currLine = lines[i]
		}
	}
	return suitableLines
}

// CountOfLines возвращает слайс булов, где true обозначает строки, для 
// которых указано количество в соответствующем элементе общего слайса строк.
// lines - строки после применения ignoreFuncs, 
// linesWithCounts - указатель на ориг строки, к ним прилепим кол-во.
func CountOfLines(lines []string, linesWithCounts *[]string) []bool { // -c
	if len(lines) == 0 {
		return []bool{}
	}

	suitableLines := make([]bool, len(lines))
	count := 1
	currLineIdx := 0
	suitableLines[0] = true
	for i := 1; i < len(lines); i++ {
		if lines[i] == lines[currLineIdx] { // spetial situation
			count++
			continue
		}
		(*linesWithCounts)[currLineIdx] = strconv.Itoa(count) + " " + (*linesWithCounts)[currLineIdx]
		suitableLines[i] = true
		currLineIdx = i
		count = 1
	}
	(*linesWithCounts)[currLineIdx] = strconv.Itoa(count) + " " + (*linesWithCounts)[currLineIdx]
	return suitableLines
}

// RepeatingLines возвращает слайс булов, где true обозначает строки, которые повторяются в данном слайсе строк.
func RepeatingLines(lines []string) []bool { // -d
	if len(lines) == 0 {
		return []bool{}
	}

	suitableLines := make([]bool, len(lines))
	var currLine string
	currIdx := 0

	for i := 0; i < len(lines); i++ {
		if lines[i] != currLine {
			currLine = lines[i]
			currIdx = i
		} else {
			suitableLines[currIdx] = true
		}
	}

	return suitableLines
}

// NonRepeatingLines возвращает слайс булов, где true обозначает строки, которые не повторяются в данном слайсе строк.
func NonRepeatingLines(lines []string) []bool { // -u
	if len(lines) == 0 {
		return []bool{}
	}

	suitableLines := make([]bool, len(lines))

	if len(lines) > 1 {
		if lines[0] != lines[1] {
			suitableLines[0] = true
		}

		for i := 1; i < len(lines)-1; i++ {
			if lines[i] != lines[i-1] && lines[i] != lines[i+1] {
				suitableLines[i] = true
			}
		}

		if lines[len(lines)-1] != lines[len(lines)-2] {
			suitableLines[len(lines)-1] = true
		}
	} else if len(lines) == 1 {
		suitableLines[0] = true
	}
	return suitableLines
}
