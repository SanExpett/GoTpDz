package inout

import (
	"bufio"
	"fmt"
	"os"
)

func FromInputToSlice(inputFileName string) []string {
	if inputFileName != "" {
		return fromFileToSlice(inputFileName)
	} else {
		return fromStdinToSlice()
	}
}

func fromFileToSlice(inputFileName string) []string {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer inputFile.Close()

	lines := []string{}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return lines
}

func fromStdinToSlice() []string {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return lines
}
