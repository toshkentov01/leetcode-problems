package main

import (
	"fmt"
	"plus_one/source"
)

func main() {
	result := source.PlusOne([]int{9}) 
	fmt.Println(result) // [1 0]

	result2 := source.PlusOne([]int{1, 2, 9, 9, 9})
	fmt.Println(result2) // [1 3 0 0 0]
}