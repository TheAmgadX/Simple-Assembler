package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Prepare Instruction Set for using...
type instruction struct {
	instType string
	bCode    []string
	hexCode  []string
	decCode  []int
}

func instructionsSet() map[string]instruction {
	instructionsMap := map[string]instruction{
		"AND": {
			instType: "mri", bCode: []string{"0000", "1000"},
			hexCode: []string{"0", "8"}, decCode: []int{0, 8},
		},
		"ADD": {
			instType: "mri", bCode: []string{"0001", "1001"},
			hexCode: []string{"1", "9"}, decCode: []int{1, 9},
		},
		"LDA": {
			instType: "mri", bCode: []string{"0010", "1010"},
			hexCode: []string{"2", "A"}, decCode: []int{2, 10},
		},
		"STA": {
			instType: "mri", bCode: []string{"0011", "1011"},
			hexCode: []string{"3", "B"}, decCode: []int{3, 11},
		},
		"BUN": {
			instType: "mri", bCode: []string{"0100", "1100"},
			hexCode: []string{"4", "C"}, decCode: []int{4, 12},
		},
		"BSA": {
			instType: "mri", bCode: []string{"0101", "1101"},
			hexCode: []string{"5", "D"}, decCode: []int{5, 13},
		},
		"ISZ": {
			instType: "mri", bCode: []string{"0110", "1110"},
			hexCode: []string{"6", "E"}, decCode: []int{6, 14},
		},
		"CLA": {
			instType: "rri", bCode: []string{"0111100000000000"},
			hexCode: []string{"7800"},
		},
		"CLE": {
			instType: "rri", bCode: []string{"0111010000000000"},
			hexCode: []string{"7400"},
		},
		"CMA": {
			instType: "rri", bCode: []string{"0111001000000000"},
			hexCode: []string{"7200"},
		},
		"CME": {
			instType: "rri", bCode: []string{"0111000100000000"},
			hexCode: []string{"7100"},
		},
		"CIR": {
			instType: "rri", bCode: []string{"0111000010000000"},
			hexCode: []string{"7080"},
		},
		"CIL": {
			instType: "rri", bCode: []string{"0111000001000000"},
			hexCode: []string{"7040"},
		},
		"INC": {
			instType: "rri", bCode: []string{"0111000000100000"},
			hexCode: []string{"7020"},
		},
		"SPA": {
			instType: "rri", bCode: []string{"0111000000010000"},
			hexCode: []string{"7010"},
		},
		"SNA": {
			instType: "rri", bCode: []string{"0111000000001000"},
			hexCode: []string{"7008"},
		},
		"SZA": {
			instType: "rri", bCode: []string{"0111000000000100"},
			hexCode: []string{"7004"},
		},
		"SZE": {
			instType: "rri", bCode: []string{"0111000000000010"},
			hexCode: []string{"7002"},
		},
		"HLT": {
			instType: "rri", bCode: []string{"0111000000000001"},
			hexCode: []string{"7001"},
		},
		"INP": {
			instType: "ioi", bCode: []string{"1111100000000000"},
			hexCode: []string{"F800"},
		},
		"OUT": {
			instType: "ioi", bCode: []string{"1111010000000000"},
			hexCode: []string{"F400"},
		},
		"SKI": {
			instType: "ioi", bCode: []string{"1111001000000000"},
			hexCode: []string{"F200"},
		},
		"SKO": {
			instType: "ioi", bCode: []string{"1111000100000000"},
			hexCode: []string{"F100"},
		},
		"ION": {
			instType: "ioi", bCode: []string{"1111000010000000"},
			hexCode: []string{"F080"},
		},
		"IOF": {
			instType: "ioi", bCode: []string{"1111000001000000"},
			hexCode: []string{"F040"},
		},
		"ORG": {
			instType: "pseudo",
		},
		"END": {
			instType: "pseudo",
		},
		"DEC": {
			instType: "pseudo",
		},
		"HEX": {
			instType: "pseudo",
		},
	}

	return instructionsMap
}

// Functions To get the code lines in slice from the txt file

func readCodeLinesFromFile() []string {
	// open the file
	file, err := os.Open("instructions.txt")

	if err != nil {
		fmt.Println("error: ", err)
		return nil
	}

	// read from file
	reader := bufio.NewReader(file)
	code := make([]string, 0)
	defer file.Close() // Ensure the file is closed

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			// End of file so break the loop
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err) // for other errors
			return nil
		}

		line = strings.TrimSpace(line)
		if line != "" {
			code = append(code, line)
		}
	}
	return code
}

// First Pass
/*
	To generate a table that correlates all user defined
	(address) symbols with their binary equivalent value
*/

// CustomSplit splits a string based on a delimiter ignoring repeated delimiters.
func CustomSplit(input, delim string) []string {
	// Replace multiple occurrences of the delimiter with a single one
	normalized := input
	for strings.Contains(normalized, delim+delim) {
		normalized = strings.ReplaceAll(normalized, delim+delim, delim)
	}

	// Trim leading and tailing delimiters
	normalized = strings.Trim(normalized, delim)

	return strings.Split(normalized, delim)
}

type label struct {
	lc            string
	valueInBinary string
	valueInDec    int
}

func ConvertLabelValue(line []string) (valDec int64, valBinary string, err error) {
	// if the pseudo instruction `DEC`
	if line[1] == "DEC" {
		// converting the value of the label to decimal
		valDec, err = strconv.ParseInt(line[2], 10, 64)
		// if the pseudo instruction `HEX`
	} else if line[1] == "HEX" {
		// converting the value of the label to decimal
		valDec, err = strconv.ParseInt(line[2], 16, 64)
	}

	if err != nil {
		return 0, "", err
	}

	// converting the value of the label to binary
	if valDec >= 0 {
		// Convert positive decimal to binary and pad with leading zeros to make it 16 bits
		valBinary = fmt.Sprintf("%016b", valDec)
	} else {
		// For negative numbers use two's complement
		// two's complement: 65536 + decimalNumber
		twoComplement := (1 << 16) + valDec
		valBinary = fmt.Sprintf("%016b", twoComplement)
	}
	return valDec, valBinary, err
}

func incrementLC(numStr string) (string, error) {
	// convert the string to int hex
	result, err := strconv.ParseInt(numStr, 16, 0)
	if err != nil {
		fmt.Println(err)
	}

	// increment and return hex string
	return fmt.Sprintf("%x", result+1), err
}

func decrementLC(numStr string) (string, error) {
	// convert the string to int hex
	result, err := strconv.ParseInt(numStr, 16, 0)
	if err != nil {
		fmt.Println(err)
	}

	// increment and return hex string
	return fmt.Sprintf("%x", result-1), err
}

func performFirstPass(code []string) (symbolTable map[string]label, LC string) {
	symbolTable = make(map[string]label)
	lc := ""

	// we are in the first line so we want to store value of ORG in lc
	if org := strings.Index(code[0], "ORG"); org != -1 {

		// splite the code line into slice to get the value of the ORG instruction
		lineComponents := CustomSplit(code[0], " ")

		lc = lineComponents[1]
		LC = lc

		var err error

		lc, err = decrementLC(lc) // to make it start from this number

		if err != nil {
			fmt.Println(err)
		}

	}

	// end of the program.
	if end := strings.Index(code[0], "END"); end != -1 {
		return
	}

	for _, line := range code {
		idx := strings.Index(line, ",")

		// we don't found label
		if idx == -1 {
			var err error

			lc, err = incrementLC(lc)

			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		// we found a label im the line then we save its value in the symbolTable
		lineComponents := CustomSplit(line, " ")

		symbol := lineComponents[0]
		symbol = symbol[:len(symbol)-1] // to avoid saving , in the variable

		valDec, valBinary, err := ConvertLabelValue(lineComponents)

		if err != nil {
			return nil, "-1"
		}

		lbl := label{
			lc:            lc,
			valueInBinary: valBinary,
			valueInDec:    int(valDec),
		}
		symbolTable[symbol] = lbl

		var IncErr error

		lc, IncErr = incrementLC(lc)

		if IncErr != nil {
			fmt.Println(IncErr)
		}
	}

	return symbolTable, LC
}

// Second Pass --> binary translation

func isLabelLine(line string) bool {
	idx := strings.Index(line, ",")

	// we found label if idx != -1
	return idx != -1
}

// to convert the LC to Binary Address of 12 bits
func convertHexToBinary(hexStr string) (string, error) {
	// Convert the hexadecimal string to an integer
	num, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error")
		return "", err
	}

	// Convert the integer to a binary string
	binaryStr := strconv.FormatInt(num, 2)

	// Ensure it's 12 bits
	if len(binaryStr) < 12 {
		binaryStr = strings.Repeat("0", 12-len(binaryStr)) + binaryStr
	}

	return binaryStr, nil
}

func convertMRIToBinary(line []string, instSet map[string]instruction,
	table map[string]label) (code string) {

	bits12To14 := ""
	bits0To11 := ""

	// indirect addressing
	if line[2] == "I" {

		// set the opCode indirect addressing
		bits12To14 = instSet[line[0]].bCode[1]

		// get the effictive address
		bits0To11 = table[line[1]].valueInBinary[4:]

		// direnct addressing
	} else {
		// set opCode direct addressing
		bits12To14 = instSet[line[0]].bCode[0]
		b, err := convertHexToBinary(table[line[1]].lc)
		if err != nil {
			return ""
		}
		bits0To11 = b
	}

	code = bits12To14 + bits0To11
	return code
}

func performSecondPass(table map[string]label,
	instructionsSet map[string]instruction, code []string, lc string) map[string]string {

	machineCode := make(map[string]string)

	for i, line := range code {

		// avoid the first line
		if i == 0 {
			continue
		}

		// end of the program
		if end := strings.Index(code[0], "END"); end != -1 {
			return machineCode
		}

		// convert the line to code
		lineCode := CustomSplit(line, " ")

		LC := lc

		keyForLabel := lineCode[0]

		if isLabelLine(line) {
			machineCode[LC] = table[keyForLabel[:len(keyForLabel)-1]].valueInBinary

			var err error

			lc, err = incrementLC(lc)

			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		// check instruction type
		switch instructionsSet[lineCode[0]].instType {
		case "mri":
			machineCode[LC] = convertMRIToBinary(lineCode, instructionsSet, table)
		case "ioi":
			machineCode[LC] = instructionsSet[lineCode[0]].bCode[0]
		case "rri":
			machineCode[LC] = instructionsSet[lineCode[0]].bCode[0]
		case "pseudo":
			if lineCode[0] == "END" {
				return machineCode
			}
		default:
			machineCode["Error"] = "Instruction Not Found"
		}

		var err error

		lc, err = incrementLC(lc)

		if err != nil {
			fmt.Println(err)
		}
	}

	return machineCode
}

func sortSliceByHex(slice []string) ([]string, error) {
	sort.Slice(slice, func(i, j int) bool {
		// Convert strings to integers (but dealing with it as hexadecimal values)
		val1, err1 := strconv.ParseInt(slice[i], 16, 64)
		val2, err2 := strconv.ParseInt(slice[j], 16, 64)

		// Handle conversion errors
		if err1 != nil || err2 != nil {
			return false
		}

		// Compare the values
		return val1 < val2
	})

	// Validate all entries
	for _, s := range slice {
		_, err := strconv.ParseInt(s, 16, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid hex value: %s", s)
		}
	}

	return slice, nil
}

func main() {
	code := readCodeLinesFromFile()

	table, LC := performFirstPass(code)

	// Symbol Table
	fmt.Println("-------------------- Symbol Table -----------------------")
	for key, val := range table {
		fmt.Println(key, " : ", val.lc)
	}
	fmt.Println("---------------------------------------------------------")

	instructionsSet := instructionsSet()

	machineCodeMap := performSecondPass(table, instructionsSet, code, LC)

	// to sort the machine code output.
	keys := make([]string, 0, len(machineCodeMap))

	for k := range machineCodeMap {
		keys = append(keys, k)
	}

	keys, _ = sortSliceByHex(keys)

	fmt.Println("-------------------- Machine Code -----------------------")
	for _, key := range keys {
		fmt.Println(key, " : ", machineCodeMap[key])
	}
	fmt.Println("---------------------------------------------------------")
}
