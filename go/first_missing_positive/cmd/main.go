package main

import (
	"first_missing_positive/source"
	"fmt"
)

func main() {
	result := source.FirstMissingPositive([]int{1, 5, 6, 7})
	fmt.Println(result) // 2

	result = source.FirstMissingPositive([]int{3, 4, 9, 7})
	fmt.Println(result) // 1

	result = source.FirstMissingPositive([]int{9, 1, 3})
	fmt.Println(result) // 2
}