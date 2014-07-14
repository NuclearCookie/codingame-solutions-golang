package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	var n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)
	horseStrenghtArr := make([]int, n)
	var outputString string
	for i := 0; i < n; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &horseStrenghtArr[i])
	}
	sort.Ints(horseStrenghtArr)
	if n > 1 {
		var diff int
		//we check i and i -1 so loop until > 1
		for i := n - 1; i > 1; i-- {
			newDiff := horseStrenghtArr[i] - horseStrenghtArr[i-1]
			if i == n-1 {
				diff = newDiff
			} else if newDiff < diff {
				diff = newDiff
			}
		}
		outputString = fmt.Sprintf("%d", diff)
	} else {
		outputString = "0"
		log.Println("no entries in data set :(")
	}
	fmt.Print(outputString)

}
