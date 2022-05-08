package main

import (
	"fmt"
	"insert_position/source"
)

func main() {
	result := source.SearchInsert([]int{1, 2, 4, 6}, 3)
	fmt.Println(result) // 2

	result = source.SearchInsert([]int{1, 2, 3}, 4)
	fmt.Println(result) // 3
}
