package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//parse the input
	var n int
	fmt.Scanln(&n)
	if n > 0 {
		temps := make([]int, n)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		stringSlice := strings.SplitAfter(text, " ")
		for i := range temps {
			var err error
			temps[i], err = strconv.Atoi(strings.TrimSpace(stringSlice[i]))
			if err != nil {
				log.Printf("Error converting string \"%s\"to integer. Check input file\n", strings.TrimSpace(stringSlice[i]))
			}
		}
		fmt.Printf("%d", GetSmallestInteger(temps))
	} else {
		log.Print("no elements to process, returning 0\n")
		fmt.Printf("0")
	}

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
		log.Print("passed slice has invalid size...\n")
	}
	return 0
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
