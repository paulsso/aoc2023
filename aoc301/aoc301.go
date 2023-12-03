package main

// part 2
// given r=(x1...xn), g=(y1...yn), b=(z1...zn) for a game
// what is the smallest number of cube that game could have been played with

import (
	"fmt"
	"os"
	s "strings"
)

var p = fmt.Println

func check(e error) {
	if e != nil {
		panic(e)
	}
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
	file, err := os.ReadFile("test.txt")
	check(err)
	inputAsListOfStrings := splitInputToArrayOfString(file)

	var sum int = 0

	// for i := 0; i < len(inputAsListOfStrings); i++ {

	// }
	p(sum)
}
