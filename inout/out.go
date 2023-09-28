package inout

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func FromSliceToOutput(lines []string, outputFileName string) error {
	if outputFileName == "" {
		fromSliceToStdout(lines)
	} else {
		err := fromSliceToFile(lines, outputFileName)
		if err != nil {
			return fmt.Errorf("failed outputing: %w", err)
		} 
	}
	
	return nil
}

func fromSliceToFile(lines []string, outputFileName string) error {
	joinedLines := strings.Join(lines, "\n")
	err := ioutil.WriteFile(outputFileName, []byte(joinedLines), 0644)

	if err != nil {
		return fmt.Errorf("failed outputing: %w", err)
	}

	return nil
}

func fromSliceToStdout(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
