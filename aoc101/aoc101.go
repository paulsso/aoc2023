package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.ReadFile("101_test.txt")
	check(err)
	var ints_as_strings []string
	var combined_ints_as_string string
	var j, sum int = 0, 0

	for i := 0; i < len(file); i++ {
		_, e := strconv.Atoi(string(file[i]))
		if e == nil {
			j++
			ints_as_strings = append(ints_as_strings, string(file[i]))
		}
		if string(file[i]) == "\n" {
			slice_length := len(ints_as_strings)
			if slice_length == 1 {
				combined_ints_as_string = s.Join([]string{ints_as_strings[0], ints_as_strings[0]}, "")
			} else {
				combined_ints_as_string = s.Join([]string{ints_as_strings[0], ints_as_strings[slice_length-1]}, "")
			}
			num, err := strconv.Atoi(combined_ints_as_string)
			check(err)
			sum = sum + num
			ints_as_strings = nil
			j = 0
		}
	}

	// As input didnt include a linebreak on the last line, run the steps once more when the reader is at EOF
	slice_length := len(ints_as_strings)
	if slice_length == 1 {
		combined_ints_as_string = s.Join([]string{ints_as_strings[0], ints_as_strings[0]}, "")
	} else {
		combined_ints_as_string = s.Join([]string{ints_as_strings[0], ints_as_strings[slice_length-1]}, "")
	}
	num, err := strconv.Atoi(combined_ints_as_string)
	check(err)
	sum = sum + num
	ints_as_strings = nil

	fmt.Println(sum)
}
