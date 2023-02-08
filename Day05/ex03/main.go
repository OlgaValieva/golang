package main

import (
	"fmt"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap struct {
	array []Present
}

func grabPresents(items []Present, capacity int) []Present {

	values := make([][]int, len(items)+1)
	for i := range values {
		values[i] = make([]int, capacity+1)
	}
	keep := make([][]int, len(items)+1)
	for i := range keep {
		keep[i] = make([]int, capacity+1)
	}
	for i := int(0); i < capacity+1; i++ {
		values[0][i] = 0
		keep[0][i] = 0
	}
	for i := 0; i < len(items)+1; i++ {
		values[i][0] = 0
		keep[i][0] = 0
	}

	for i := 1; i <= len(items); i++ {
		for c := int(1); c <= capacity; c++ {
			itemFits := (items[i-1].Size <= c)
			if !itemFits {
				continue
			}

			maxValueAtThiscapacityacity := items[i-1].Value + values[i-1][c-items[i-1].Size]
			previousValueAtThisCapacity := values[i-1][c]

			if itemFits && (maxValueAtThiscapacityacity > previousValueAtThisCapacity) {
				values[i][c] = maxValueAtThiscapacityacity
				keep[i][c] = 1
			} else {
				values[i][c] = previousValueAtThisCapacity
				keep[i][c] = 0
			}
		}
	}

	n := len(items)
	c := capacity

	var res []Present

	for n > 0 {
		if keep[n][c] == 1 {
			c -= items[n-1].Size
			res = append(res, items[n-1])
		}
		n--
	}
	return res

}

func main() {
	array := []Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}
	fmt.Println(array)
	grabbed := grabPresents(array, 5)	
	fmt.Println(grabbed)
}