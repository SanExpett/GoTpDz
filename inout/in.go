package inout

import (
	"bufio"
	"fmt"
	"os"
)

func FromInputToSlice(inputFileName string) ([]string, error) {
	if inputFileName != "" {
		result, err := fromFileToSlice(inputFileName)
		if err != nil {
			return nil, err
		} 
		return result, nil
	} 
	result, err := FromInputToSlice(inputFileName)
	if err != nil {
		return nil, err
	} 
	return result, nil
}

func fromFileToSlice(inputFileName string) ([]string, error) {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return lines, nil
}

func fromStdinToSlice() ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
