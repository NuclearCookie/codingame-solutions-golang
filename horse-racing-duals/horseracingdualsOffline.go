package main

import (
	"fmt"
	"github.com/nuclearcookie/cgreader"
	"sort"
)

func main() {
	cgreader.RunStaticPrograms(
		cgreader.GetFileList("input/input%d.txt", 3),
		cgreader.GetFileList("output/output%d.txt", 3),
		true,
		func(input <-chan string, output chan string) {
			var n int
			fmt.Sscanf(<-input, "%d", &n)
			horseStrenghtArr := make([]int, n)
			var outputString string
			for i := 0; i < n; i++ {
				fmt.Sscanf(<-input, "%d", &horseStrenghtArr[i])
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
					} else if newDiff == diff {
						cgreader.Traceln("2 horses with the same strenght diff found!")
					}
				}
				outputString = fmt.Sprintf("%d", diff)
			} else {
				outputString = "0"
				cgreader.Traceln("no entries in data set :(")
			}
			output <- outputString
		})
}
