package main

// part 1
// problem synposis
// given r:12, g:13, b:14
// determine the subset of possible combinations r=x, g=y, b=z
// out of the given set of combinations (input)

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

var p = fmt.Println

type draw struct {
	color string
	num   int
}

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

func gameOkCheck(game string) bool {
	_game := s.Split(game, ":")
	_draws := s.Split(_game[1], ";")

	redCap := 12
	greenCap := 13
	blueCap := 14

	for i := 0; i < len(_draws); i++ {
		_d := s.SplitAfter(_draws[i], ",")
		for j := 0; j < len(_d); j++ {
			__d := s.Replace(_d[j], " ", "", 1)
			if s.Contains(__d, "blue") {
				digit := s.Split(__d, " ")
				numBlues, _ := strconv.Atoi(digit[0])
				if blueCap < numBlues {
					return false
				}
			} else if s.Contains(__d, "red") {
				digit := s.Split(__d, " ")
				numReds, _ := strconv.Atoi(digit[0])
				if redCap < numReds {
					return false
				}
			} else if s.Contains(__d, "green") {
				digit := s.Split(__d, " ")
				numGreens, _ := strconv.Atoi(digit[0])
				if greenCap < numGreens {
					return false
				}
			}
		}
	}
	return true
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
	file, err := os.ReadFile("201_input.txt")
	check(err)
	inputAsListOfStrings := splitInputToArrayOfString(file)
	var sum int
	for i := 0; i < len(inputAsListOfStrings); i++ {
		gameOk := gameOkCheck(inputAsListOfStrings[i])
		if gameOk {
			sum += i + 1
		}
	}
	p(sum)
}
