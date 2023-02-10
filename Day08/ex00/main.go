package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty slice")
	}
	if idx < 0 {
		return 0, errors.New("negative index")
	}
	if len(arr) <= idx {
		return 0, errors.New("index is out of bounds")
	}
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + uintptr(idx)*unsafe.Sizeof(arr[0]))), nil
}

func main() {
	var err error
	var res int
	var idx = 0
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr: ", arr)

	res, err = getElement(arr, idx)

	fmt.Printf("idx = %d, res = %d error: %v\n", idx, res, err)

	idx = -1
	res, err = getElement(arr, idx)
	fmt.Printf("idx = %d, res = %d error: %v\n", idx, res, err)

	idx = 10
	res, err = getElement(arr, idx)
	fmt.Printf("idx = %d, res = %d error: %v\n\n", idx, res, err)

	empty := []int{}
	fmt.Println("empty arr: ", empty)
	idx = 7
	res, err = getElement(empty, idx)
	fmt.Printf("idx = %d, res = %d error: %v\n", idx, res, err)
}
