package main

// part 2
// given r=(x1...xn), g=(y1...yn), b=(z1...zn) for a game
// what is the smallest number of cube that game could have been played with

import (
	"fmt"
	"os"
	"regexp"
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

func getNumberOfCols(file []byte) int {
	var j int = 0
	for i := 0; i < len(file); i++ {
		if string(file[i]) == "\n" {
			break
		}
		j++
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

func checkOneIndexRange(number_irange []int, symbols_irange [][]int) bool {
	for i := 0; i < len(symbols_irange); i++ {
		if symbols_irange[i][0] == number_irange[0]-1 {
			return true
		} else if symbols_irange[i][1] == number_irange[1]+1 {
			return true
			// check if there is a symbol with index range inside number_irange[0]-1:number_irange[1]+1
		} else if (symbols_irange[i][0] >= number_irange[0]-1) && (symbols_irange[i][1] <= number_irange[1]+1) {
			return true
		}
	}
	return false
}

func returnValidIndex(main string, below string, above string) [][]int {
	symbol, _ := regexp.Compile("([^\\.0-9])")
	number, _ := regexp.Compile("([0-9]{1,3})")

	var valid_indeces [][]int

	number_index := number.FindAllStringIndex(main, -1)
	symbols_same_row := symbol.FindAllStringIndex(main, -1)
	symbols_below_row := symbol.FindAllStringIndex(below, -1)
	symbols_above_row := symbol.FindAllStringIndex(above, -1)

	for i := 0; i < len(number_index); i++ {
		if checkOneIndexRange(number_index[i], symbols_same_row) {
			valid_indeces = append(valid_indeces, number_index[i])
		} else if checkOneIndexRange(number_index[i], symbols_above_row) {
			valid_indeces = append(valid_indeces, number_index[i])
		} else if checkOneIndexRange(number_index[i], symbols_below_row) {
			valid_indeces = append(valid_indeces, number_index[i])
		}
	}
	return valid_indeces
}

func returnValidIndexEnd(main string, next string) [][]int {
	symbol, _ := regexp.Compile("([^\\.0-9])")
	number, _ := regexp.Compile("([0-9]{1,3})")

	var valid_indeces [][]int

	number_index := number.FindAllStringIndex(main, -1)
	symbols_same_row := symbol.FindAllStringIndex(main, -1)
	symbols_next_row := symbol.FindAllStringIndex(next, -1)

	for i := 0; i < len(number_index); i++ {
		if checkOneIndexRange(number_index[i], symbols_same_row) {
			valid_indeces = append(valid_indeces, number_index[i])
		} else if checkOneIndexRange(number_index[i], symbols_next_row) {
			valid_indeces = append(valid_indeces, number_index[i])
		}
	}
	return valid_indeces
}

func main() {
	file, err := os.ReadFile("real.txt")
	check(err)
	input := splitInputToArrayOfString(file)
	var sum int = 0
	for i := 1; i < len(input)-1; i++ {
		valid_index := returnValidIndex(input[i], input[i+1], input[i-1])
		for j := 0; j < len(valid_index); j++ {
			str := input[i]
			p(str[valid_index[j][0]:valid_index[j][1]])
			num, _ := strconv.Atoi(str[valid_index[j][0]:valid_index[j][1]])
			sum += num
		}
	}

	valid_index := returnValidIndexEnd(input[0], input[1])
	for j := 0; j < len(valid_index); j++ {
		str := input[0]
		p(str[valid_index[j][0]:valid_index[j][1]])
		num, _ := strconv.Atoi(str[valid_index[j][0]:valid_index[j][1]])
		sum += num
	}

	valid_index = returnValidIndexEnd(input[len(input)-1], input[len(input)-2])
	for j := 0; j < len(valid_index); j++ {
		str := input[len(input)-1]
		p(str[valid_index[j][0]:valid_index[j][1]])
		num, _ := strconv.Atoi(str[valid_index[j][0]:valid_index[j][1]])
		sum += num
	}

	p(sum)
}
