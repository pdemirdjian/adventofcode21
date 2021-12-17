package day10

import "adventofcode21/utility"

func Day10(inputFileName string) int {
	lines := utility.GetArrayOfInputFile(inputFileName)
	badChars := processLines(lines)
	sum := 0
	for _, badChar := range badChars {
		switch badChar {
		case ")":
			sum = sum + 3
		case "]":
			sum = sum + 57
		case "}":
			sum = sum + 1197
		case ">":
			sum = sum + 25137
		}
	}
	return sum
}

func processLines(lines []string) []string {
	var badChars []string
	for _, line := range lines {
		badChars = append(badChars, processLine(line))
	}
	return badChars
}

func processLine(line string) string {
	var brackets []string
	for _, char := range line {
		charString := string(char)
		if charString == "<" || charString == "[" || charString == "{" || charString == "(" {
			brackets = append(brackets, charString)
		} else {
			lastIndex := len(brackets) - 1
			lastElement := brackets[lastIndex]
			switch charString {
			case ">":
				if lastElement != "<" {
					return charString
				} else {
					brackets = utility.PopLastElementStringArr(brackets)
				}
			case "]":
				if lastElement != "[" {
					return charString
				} else {
					brackets = utility.PopLastElementStringArr(brackets)
				}
			case "}":
				if lastElement != "{" {
					return charString
				} else {
					brackets = utility.PopLastElementStringArr(brackets)
				}
			case ")":
				if lastElement != "(" {
					return charString
				} else {
					brackets = utility.PopLastElementStringArr(brackets)
				}
			}
		}
	}
	return ""
}
