package utility

import (
	"bufio"
	"log"
	"os"
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
