package inout

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func FromSliceToOutput(lines []string, outputFileName string) {
	if outputFileName != "" {
		fromSliceToFile(lines, outputFileName)
	} else {
		fromSliceToStdout(lines)
	}
}

func fromSliceToFile(lines []string, outputFileName string) {
	joinedLines := strings.Join(lines, "\n")
	err := ioutil.WriteFile(outputFileName, []byte(joinedLines), 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func fromSliceToStdout(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
