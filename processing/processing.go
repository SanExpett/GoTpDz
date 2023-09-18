package processing

import (
	"flag"
	"fmt"

	"github.com/SanExpett/TpGoDz/functions"
	"github.com/SanExpett/TpGoDz/inout"
)

func initFlags(cFlag, dFlag, uFlag, iFlag *bool, fFlag, sFlag *int) {
	flag.BoolVar(cFlag, "c", false, "Count before each string")
	flag.BoolVar(dFlag, "d", false, "Only reapeting")
	flag.BoolVar(uFlag, "u", false, "Only non-reapeting")
	flag.IntVar(fFlag, "f", 0, "Ignore first num fields")
	flag.IntVar(sFlag, "s", 0, "Ignore first num chars")
	flag.BoolVar(iFlag, "i", false, "Ignore register")
}

// func getResult(lines []string, suitableLines []bool) []string {
// 	var result []string
// 	for i, val := range suitableLines {
// 		if val {
// 			result = append(result, lines[i])
// 		}
// 	}
// 	return result 
// }

//uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]
func executeCommands(cFlag, dFlag, uFlag, iFlag bool, fFlag, sFlag int, lines *[]string) {
	linesCopy := make([]string, len(*lines))
	copy(linesCopy, *lines)

	suitableLines := make([]bool, len(linesCopy))
	// БЛОК 1
	if iFlag {
		functions.IgnoreRegister(&linesCopy) 
	} 
	if fFlag > 0 {
		functions.IgnoreNFields(&linesCopy, fFlag)
	}
	if sFlag > 0 {
		functions.IgnoreNSymbols(&linesCopy, sFlag)
	}

	// Блок 2
	if (cFlag != dFlag) != uFlag { // костыльный xor
		if cFlag {
			suitableLines = functions.CountOfLines(linesCopy, lines)
		}	
		if dFlag {
			suitableLines = functions.RepeatingLines(linesCopy)
		}
		if uFlag {
			suitableLines = functions.NonRepeatingLines(linesCopy)
		}
	} else if !(cFlag || dFlag || uFlag) {
		suitableLines = functions.UniqLines(linesCopy)
	} else {
		fmt.Println("Формат команды: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}

	*lines, _ = functions.GetResult(*lines, suitableLines)
}

func ParseCommandLine() {
	var cFlag, dFlag, uFlag, iFlag bool
	var fFlag, sFlag int
	initFlags(&cFlag, &dFlag, &uFlag, &iFlag, &fFlag, &sFlag)
	flag.Parse()

	inputFileName := flag.Arg(0)
	outputFileName := flag.Arg(1)
	
	lines := inout.FromInputToSlice(inputFileName)
	executeCommands(cFlag, dFlag, uFlag, iFlag, fFlag, sFlag, &lines)
	inout.FromSliceToOutput(lines, outputFileName)
}