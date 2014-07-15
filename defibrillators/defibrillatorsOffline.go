package main

import (
	"fmt"
	"github.com/nuclearcookie/cgreader"
	"log"
	"math"
	"strconv"
	"strings"
)

type defibrillatorInfo struct {
	ID        int
	Name      string
	Adress    string
	Phone     string
	Longitude float64
	Lattitude float64
}

func main() {
	cgreader.RunStaticPrograms(
		cgreader.GetFileList("input/input%d.txt", 4),
		cgreader.GetFileList("output/output%d.txt", 4),
		true,
		func(input <-chan string, output chan string) {
			//parse input
			longitude := AsciiFloatToFloat(<-input)
			lattitude := AsciiFloatToFloat(<-input)
			var n int
			fmt.Sscanln(<-input, &n)
			defibrillatorList := make([]defibrillatorInfo, n)
			for i := 0; i < n; i++ {
				defibrillatorList[i] = GetDefibrillatorInfoFromString(<-input)
			}
			//perform logic
			outputIndex := 0
			distance := GetDistanceBetweenHumanAndAID(longitude, lattitude, defibrillatorList[0])
			for i := 1; i < n; i++ {
				newDistance := GetDistanceBetweenHumanAndAID(longitude, lattitude, defibrillatorList[i])
				if newDistance < distance {
					outputIndex = i
					distance = newDistance
				}
			}
			//output
			output <- defibrillatorList[outputIndex].Name
		})
}

func GetDefibrillatorInfoFromString(input string) (output defibrillatorInfo) {
	info := strings.Split(input, ";")
	var err error
	output.ID, err = strconv.Atoi(info[0])
	output.Name = info[1]
	output.Adress = info[2]
	output.Phone = info[3]
	output.Longitude = AsciiFloatToFloat(info[4])
	output.Lattitude = AsciiFloatToFloat(info[5])
	if err != nil {
		log.Fatal(err)
	}
	return
}

func AsciiFloatToFloat(input string) (output float64) {
	input = strings.Replace(input, ",", ".", -1)
	fmt.Sscanf(input, "%f", &output)
	return
}

func GetDistanceBetweenHumanAndAID(longitude, lattitude float64, info defibrillatorInfo) float64 {
	x := (info.Longitude - longitude) * math.Cos((lattitude+info.Lattitude)/2.0)
	y := info.Lattitude - lattitude
	d := math.Sqrt(x*x+y*y) * 6371
	return d
}
