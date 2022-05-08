package main

import (
	"fmt"
	"single_number/source"
)

func main() {
	result := source.SingleNumber([]int{4, 1, 1, 2, 3, 5, 3, 5, 2})
	fmt.Println(result) // 4
}
