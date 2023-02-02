package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"	
	"sync"
	"strings"
)

var lines = false
var chars = false
var words = false

func checkFlag() bool {
	flag.BoolVar(&lines, "l", false, "lines")
	flag.BoolVar(&chars, "m", false, "characters")
	flag.BoolVar(&words, "w", false, "words")
	flag.Parse()
	if !lines && !chars && !words {
		words = true
	}
	if (lines && !chars && !words) || (chars && !lines && !words) || (words && !lines && !chars)   {
		return true
	}	
	return false
}

func main() {
	if !checkFlag() {
		fmt.Println("only one flag can be specified at a time")
		return
	}
	var wg sync.WaitGroup
	for _, file := range flag.Args() {
		wg.Add(1)
		file := file
		go func() {
			if lines {
				fmt.Printf("%d\t%s\n", cl(file), file)
			} else if words {
				fmt.Printf("%d\t%s\n", cw(file), file)
			} else if chars {
				fmt.Printf("%d\t%s\n", cm(file), file)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func cl(s string) int {
	var lines int
	file, _ := os.ReadFile(s)
	for range strings.Split(strings.TrimSuffix(string(file), "\n"), "\n") {
		lines++
	}
	return lines
}

func cw(s string) int {
	var word int
	fileHandle, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer func(fileHandle *os.File) {
		err := fileHandle.Close()
		if err != nil {}
	}(fileHandle)

	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanWords)
	for fileScanner.Scan() {
		word++
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}
	return word
}

func cm(s string) int {
	file, _ := os.ReadFile(s)
	return len(file)
}