package day03

import (
	"adventofcode21/utility"
	"log"
	"strconv"
)

func Day03(fileInputPath string) int {
	inputs := utility.GetArrayOfInputFile(fileInputPath)

	mostBits := make([]int, len(inputs[0]))

	for _, input := range inputs {
		for i := 0; i < len(input); i++ {
			char := rune(input[i])
			if char == '1' {
				mostBits[i] = mostBits[i] + 1
			} else {
				mostBits[i] = mostBits[i] - 1
			}
		}
	}
	gammaValBits := ""
	epsilonValBits := ""
	for _, bit := range mostBits {
		if bit > 0 {
			gammaValBits = gammaValBits + "1"
			epsilonValBits = epsilonValBits + "0"
		} else {
			gammaValBits = gammaValBits + "0"
			epsilonValBits = epsilonValBits + "1"
		}
	}
	gammaVal, err := strconv.ParseInt(gammaValBits, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}
	epsilonVal, err := strconv.ParseInt(epsilonValBits, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return int(gammaVal) * int(epsilonVal)
}

func Day0302(fileInputPath string) int {
	inputs := utility.GetArrayOfInputFile(fileInputPath)

	O2Inputs := inputs
	CO2Inputs := inputs

	for i := 0; i < len(inputs[0]); i++ {
		mostBit := findMostBitsInPos(O2Inputs, i)
		if mostBit > 0 {
			O2Inputs = trimInputs(O2Inputs, i, 1)
		} else if mostBit < 0 {
			O2Inputs = trimInputs(O2Inputs, i, -1)
		} else {
			O2Inputs = trimInputs(O2Inputs, i, 1)
		}
	}

	O2Reading, err := strconv.ParseInt(O2Inputs[0], 2, 64)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < len(inputs[0]); i++ {
		mostBit := findMostBitsInPos(CO2Inputs, i)
		if mostBit > 0 {
			CO2Inputs = trimInputs(CO2Inputs, i, -1)
		} else if mostBit < 0 {
			CO2Inputs = trimInputs(CO2Inputs, i, 1)
		} else {
			CO2Inputs = trimInputs(CO2Inputs, i, -1)
		}
	}
	CO2Reading, err := strconv.ParseInt(CO2Inputs[0], 2, 64)
	if err != nil {
		log.Fatalln(err)
	}

	return int(O2Reading * CO2Reading)
}

func trimInputs(inputs []string, pos int, direction int) []string {
	if len(inputs) == 1 {
		return inputs
	}
	returns := []string{}
	bit := ""
	if direction >= 0 {
		bit = "1"
	} else {
		bit = "0"
	}
	for _, input := range inputs {
		char := string(input[pos])
		if char == bit {
			returns = append(returns, input)
		}
	}
	return returns
}

func findMostBitsInPos(inputs []string, pos int) int {

	mostBits := 0

	for _, input := range inputs {
		char := string(input[pos])
		if char == "1" {
			mostBits = mostBits + 1
		} else {
			mostBits = mostBits - 1
		}
	}

	if mostBits > 0 {
		return 1
	} else if mostBits < 0 {
		return -1
	} else {
		return 0
	}

}
