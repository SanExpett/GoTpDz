// в этом файле будет лежать код функций для флагов игнорирования чего либо (-i, -f num, -s num)

package functions

import (
	"errors"
	"strings"
)

var errIgnore = errors.New("selection funcs error")

func IgnoreRegister(lines []string) {
	for i := range lines {
		lines[i] = strings.ToLower(lines[i])
	}
}

func IgnoreNFields(lines []string, n int) error {
	if n < 0 {
		return errIgnore
	}

	for i := range lines {
		words := strings.Split(lines[i], " ")
		if n > len(words) {
			lines[i] = ""
		} else {
			lines[i] = strings.Join(words[n:], " ")
		}
	}

	return nil
}

func IgnoreNSymbols(lines []string, n int) error {
	if n < 0 {
		return errIgnore
	}

	for i := range lines {
		if n > len((lines)[i]) {
			lines[i] = ""
		} else {
			lines[i] = lines[i][n:]
		}
	}

	return nil
}
