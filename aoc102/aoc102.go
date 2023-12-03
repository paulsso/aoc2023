package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

var p = fmt.Println

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mergeStringsToInt(input string) int {
	lastIndex := len(input)
	lastChar := string(input[lastIndex-1])
	output, _ := strconv.Atoi(s.Join([]string{string(input[0]), lastChar}, ""))
	return output
}

func extractIntegersAsStringElements(input string) string {
	var output string
	for i := 0; i < len(input); i++ {
		_, err := strconv.Atoi(string(input[i]))
		if err == nil {
			output = s.Join([]string{output, string(input[i])}, "")
		}
	}
	return output
}

func applyFilter(filters [9]string, masks [9]string, input string) string {
	for j := 0; j < len(filters); j++ {
		input = s.Replace(input, filters[j], masks[j], -1)
	}
	return input
}

func getNumberOfLines(file []byte) int {
	var j int = 1
	for i := 0; i < len(file); i++ {
		if string(file[i]) == "\n" {
			j++
		}
	}
	return j
}

func splitInputToArrayOfString(file []byte) []string {
	var nLines int = getNumberOfLines(file)
	var list []string
	var window int = 0
	for i := 0; i < nLines; i++ {
		var slice string
		for j := window; j < len(file); j++ {
			if string(file[j]) != "\n" {
				slice = s.Join([]string{slice, string(file[j])}, "")
			} else {
				list = append(list, slice)
				window = j + 1
				break
			}
		}
		// The end of the input does not have a new line,
		// therefore append the last slice manually
		if i == nLines-1 {
			list = append(list, slice)
		}
	}
	return list
}

func main() {
	file, err := os.ReadFile("102_input.txt")
	check(err)

	filters := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	masks := [9]string{"o1e", "t2o", "th3ee", "f4ur", "f5ve", "s6x", "se7en", "ei8ht", "n9ne"}

	var sum int = 0

	inputAsListOfStrings := splitInputToArrayOfString(file)
	for i := 0; i < len(inputAsListOfStrings); i++ {
		filtered := applyFilter(filters, masks, inputAsListOfStrings[i])
		extracted := extractIntegersAsStringElements(filtered)
		merged := mergeStringsToInt(extracted)
		sum += merged
	}
	p(sum)
}
