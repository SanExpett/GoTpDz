package processing

import (
	"errors"
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

//uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]
func executeCommands(cFlag, dFlag, uFlag, iFlag bool, fFlag, sFlag int, lines *[]string) error {
	linesCopy := make([]string, len(*lines))
	copy(linesCopy, *lines)

	suitableLines := make([]bool, len(linesCopy))
	// флаги, которые говорят игнорировать регистр/ n слов / n символов (ломаем ими linesCopy)
	if iFlag {
		functions.IgnoreRegister(&linesCopy) 
	} 
	if fFlag != 0 {
		errF := functions.IgnoreNFields(&linesCopy, fFlag)
		if errF != nil {
			return errF
		}
	}
	if sFlag != 0 {
		errS := functions.IgnoreNSymbols(&linesCopy, sFlag)
		if errS != nil {
			return errS
		}
	}

	// флаги, которые делают выорку из строк, которые уже были изменены ignore флагами
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
		return errors.New("Формат команды: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]") 
	}

	err := functions.GetResult(lines, suitableLines)
	if err!= nil {
		return err
	}
	return nil
}

// вызывает initFlags, парсит строчку, считывает имена файлов (или пуст. стр., если их нет) вызывает FromInputToSlice, чтобы записать введенные строки в lines
// вызывает executeCommands, внутри которого меняется lines, вызывает FromSliceToOutput чтобы вывести lines в нужный вывод
func ParseCommandLine() error {
	var cFlag, dFlag, uFlag, iFlag bool
	var fFlag, sFlag int
	initFlags(&cFlag, &dFlag, &uFlag, &iFlag, &fFlag, &sFlag)
	flag.Parse()

	inputFileName := flag.Arg(0)
	outputFileName := flag.Arg(1)
	
	lines, inputErr := inout.FromInputToSlice(inputFileName)
	if inputErr != nil {
		return inputErr
	}

	execErr := executeCommands(cFlag, dFlag, uFlag, iFlag, fFlag, sFlag, &lines)
	if execErr != nil {
		return execErr
	}
	inout.FromSliceToOutput(lines, outputFileName)
	return nil
}