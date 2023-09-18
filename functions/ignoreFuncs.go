// в этом файле будет лежать код функций для флагов игнорирования чего либо (-i, -f num, -s num)

package functions

import (
	"strings"
)

func IgnoreRegister(lines *[]string) {
	for i := range *lines {
		(*lines)[i] = strings.ToLower((*lines)[i])
	}
}

func IgnoreNFields(lines *[]string, n int) {
	for i := range *lines {
		words := strings.Split((*lines)[i], " ")
		if n < 0 || n > len(words) {
			(*lines)[i] = ""
		} else {
			(*lines)[i] = strings.Join(words[n:], " ")
		}
	}
}

func IgnoreNSymbols(lines *[]string, n int) {
	for i := range *lines {
		if n < 0 || n > len( (*lines)[i] ) {
			(*lines)[i] = ""
		} else {
			(*lines)[i] = (*lines)[i][n:]
		}
	}
}
