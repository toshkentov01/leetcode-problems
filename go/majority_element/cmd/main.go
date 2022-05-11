package main

import (
	"fmt"
	"majority_element/source"
)

func main() {
	result := source.MajorityElement([]int{1, 2, 3, 3, 3})
	fmt.Println(result) // 3
}