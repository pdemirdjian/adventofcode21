package utility

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetArrayOfInputFile(inputFileName string) []string {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func GetCommaSeperatedStringsOneLine(inputFileName string) []string {
	byteContents, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	contents := string(byteContents)

	stringParts := strings.Split(contents, ",")
	return stringParts
}

func GetIntArrFromStringArr(inputs []string) []int {
	nums := []int{}
	for _, input := range inputs {
		actualInt, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalln(err)
		}
		nums = append(nums, actualInt)
	}
	return nums
}

func GetFloat64ArrFromStringArr(inputs []string) []float64 {
	nums := []float64{}
	for _, input := range inputs {
		actualInt, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatalln(err)
		}
		nums = append(nums, actualInt)
	}
	return nums
}
