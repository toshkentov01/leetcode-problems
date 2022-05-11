package main

import (
	"binary/source"
	"fmt"
)

func main() {
	result := source.Search([]int{1, 2, 3, 4}, 4)
	fmt.Println(result) // 3

	result = source.Search([]int{2, 1, 44, 2}, 1)
	fmt.Println(result) // -1
}
