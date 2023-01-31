package main

import (
	"flag"
	"math"
	"io"
	"sort"
	"fmt"
)

var meanFalse bool = false
var medianFalse bool = false
var modeFalse bool = false
var sdFalse bool = false

func checkFlags() {
	flag.BoolVar(&meanFalse, "mean", false, "mean")
	flag.BoolVar(&medianFalse, "median", false, "median")
	flag.BoolVar(&modeFalse, "mode", false, "mode")
	flag.BoolVar(&sdFalse, "sd", false, "sd")
	flag.Parse()
	if !meanFalse && !medianFalse && !modeFalse && !sdFalse {
		meanFalse = true
		medianFalse = true
		modeFalse  = true
		sdFalse = true
	}
}

func mean(slice []int) float64 {
	var mean float64
	for _, num := range slice {
		mean += float64(num)
	}
	return math.Round(mean/float64(len(slice))* 100) / 100
}

func median(slice []int) float64 {
	if len(slice)%2 == 0 {
		return mean([]int{slice[len(slice)/2-1], slice[len(slice)/2]})
	} else {
		return float64(slice[len(slice)/2])
	}
}

func mode(slice []int) (mode int) {
	countMap := make(map[int]int)
	for _, value := range slice {
		countMap[value] += 1
	}
	maxValue := 0
	for _, count := range slice {
		freq := countMap[count]
		if freq > maxValue {
			mode = count
			maxValue = freq
		}
	}
	return
}

func sd(slice []int) float64 {
	var sd float64
	var mean float64
	for _, num := range slice {
		mean += float64(num)
	}
	mean = math.Round(mean/float64(len(slice))* 100) / 100
	for _, num := range slice {
		sd += math.Pow(float64(num) - mean, 2)
	}
	return math.Round(math.Sqrt(sd/float64(len(slice))) * 100) / 100
}

func main() {
	var slice []int
	var tmp int
	checkFlags()
	inputValue, err := fmt.Scanln(&tmp)
	if inputValue == 0 {
		fmt.Println("unexpected input:letter characters or empty string\n")
		return
	}
	for err == nil {
		if tmp > -100000 && tmp < 100000 {
			slice = append(slice, tmp)
		} else {
			fmt.Println("number out of bounds")
		}
		inputValue, err = fmt.Scanln(&tmp)
	}
	if err != io.EOF {
		fmt.Println(err)
	}
	sort.Ints(slice)
	if len(slice) != 0 {
		if meanFalse {
			fmt.Println("Mean:", mean(slice))
		}
		if medianFalse {
			fmt.Println("Median:", median(slice))
		}
		if modeFalse {
			fmt.Println("Mode:", mode(slice))
		}
		if sdFalse {
			fmt.Println("SD:", sd(slice))
		}
	}
}
