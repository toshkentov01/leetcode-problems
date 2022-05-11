package main

import (
	"contains_duplicate/source"
	"fmt"
)

func main() {
	contains := source.ContainsDuplicate([]int{1, 2, 3, 1})

	if contains {
		fmt.Println("Contains")

	} else {
		fmt.Println("Not Contains")
	}
}
