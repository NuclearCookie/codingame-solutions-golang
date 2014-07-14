package main

import (
	"fmt"
	"github.com/nuclearcookie/cgreader"
	"strconv"
	"strings"
)

func main() {
	cgreader.RunStaticPrograms(
		cgreader.GetFileList("input/input%d.txt", 3),
		cgreader.GetFileList("output/output%d.txt", 3),
		true,
		func(input <-chan string, output chan string) {
			//parse the input
			var n int
			fmt.Sscanln(<-input, &n)
			if n > 0 {
				temps := make([]int, n)
				text := <-input
				stringSlice := strings.SplitAfter(text, " ")
				for i := range temps {
					var err error
					temps[i], err = strconv.Atoi(strings.TrimSpace(stringSlice[i]))
					if err != nil {
						cgreader.Tracef("Error converting string \"%s\"to integer. Check input file\n", strings.TrimSpace(stringSlice[i]))
					}
				}
				output <- fmt.Sprintf("%d", GetSmallestInteger(temps))
			} else {
				cgreader.Trace("no elements to process, returning 0\n")
				output <- "0"
			}
		})
}

//slices can be passed by value because this contains only the ptr to the internal array, and a lenght and capacity variable.
func GetSmallestInteger(slice []int) int {
	if len(slice) > 0 {
		returnValue := slice[0]
		for i := 1; i < len(slice); i++ {
			if IntAbs(slice[i]) < IntAbs(returnValue) {
				returnValue = slice[i]
			} else if IntAbs(slice[i]) == IntAbs(returnValue) && returnValue < slice[i] {
				returnValue = slice[i]
			}
		}
		return returnValue
	} else {
		cgreader.Trace("passed slice has invalid size...\n")
	}
	return 0
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
