package main

import (
	"github.com/nuclearcookie/cgreader"
)

func main() {
	cgreader.RunStaticPrograms(
		cgreader.GetFileList("input/input%d.txt", 4),
		cgreader.GetFileList("output/output%d.txt", 4),
		true,
		func(input <-chan string, output chan string) {

		})
}
