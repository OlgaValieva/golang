package main

import "fmt"

type Present struct {
	Value int
	Size  int
}

type PresentHeap struct {
	array []Present
}

func (p PresentHeap) sortirovka() {
	for i := 0; i < len(p.array); i++ {
		for j := i; j < len(p.array); j++ {
			if p.comparison(i, j) {
				p.array[i].Value, p.array[j].Value = p.array[j].Value, p.array[i].Value
				p.array[i].Size, p.array[j].Size = p.array[j].Size, p.array[i].Size
			}
		}
	}
}

func (p PresentHeap) comparison(i, j int) bool {
	if p.array[i].Value == p.array[j].Value {
		return p.array[i].Size > p.array[j].Size
	} else {
		return p.array[i].Value < p.array[j].Value
	}
}

func getNCoolestPresents(ph []Present, n int) PresentHeap {
	if n < 0  {
		fmt.Println("n is negative")
		//return PresentHeap{nil}
	}
	if  n > len(ph) {
		fmt.Println("n is larger than the size of the slice")
		//return PresentHeap{nil}
	}
	var oldPh PresentHeap = PresentHeap{ph}
	var newPh PresentHeap
	oldPh.sortirovka()

	for i, elemArr := range oldPh.array {
		if i == n {
			break
		}
		newPh.Push(elemArr)
	}
	return newPh
}

func (p *PresentHeap) Push(x any) {
	(*p).array = append((*p).array, x.(Present))
}

func main() {
	array := []Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}
	p := PresentHeap{array}
	fmt.Println(array)

	all := getNCoolestPresents(p.array, len(p.array))
	fmt.Println(all.array)

	only2 := getNCoolestPresents(p.array, 2)
	fmt.Println(only2.array)

	neg := getNCoolestPresents(p.array, len(p.array) - 10)
	fmt.Println(neg.array)

	big:= getNCoolestPresents(p.array, len(p.array) + 100)
	fmt.Println(big.array)
}