package main

// part 2
// given r=(x1...xn), g=(y1...yn), b=(z1...zn) for a game
// what is the smallest number of cube that game could have been played with

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

func getNumberOfLines(file []byte) int {
	var j int = 1
	for i := 0; i < len(file); i++ {
		if string(file[i]) == "\n" {
			j++
		}
	}
	return j
}

func minNumberOfCubesInGame(game string) [3]int {
	_game := s.Split(game, ":")
	_draws := s.Split(_game[1], ";")

	maxReds := 0
	maxGreens := 0
	maxBlues := 0

	for i := 0; i < len(_draws); i++ {
		_d := s.Split(_draws[i], ",")
		for j := 0; j < len(_d); j++ {
			__d := s.Replace(_d[j], " ", "", 1)
			if s.Contains(__d, "blue") {
				digit := s.Split(__d, " ")
				numBlues, _ := strconv.Atoi(digit[0])
				if maxBlues < numBlues {
					maxBlues = numBlues
				}
			} else if s.Contains(__d, "red") {
				digit := s.Split(__d, " ")
				numReds, _ := strconv.Atoi(digit[0])
				if maxReds < numReds {
					maxReds = numReds
				}
			} else if s.Contains(__d, "green") {
				digit := s.Split(__d, " ")
				numGreens, _ := strconv.Atoi(digit[0])
				if maxGreens < numGreens {
					maxGreens = numGreens
				}
			}
		}
	}
	return [3]int{maxReds, maxGreens, maxBlues}
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

	var sum int = 0
	for i := 0; i < len(inputAsListOfStrings); i++ {
		minCubes := minNumberOfCubesInGame(inputAsListOfStrings[i])
		sum += minCubes[0] * minCubes[1] * minCubes[2]
	}
	p(sum)
}
