package main

import (
	"fmt"
	"two_sum/source"
)

func main() {
	result := source.TwoSum([]int{1, 2, 3, 6}, 7)

	fmt.Println(result) // [3, 0]
}
