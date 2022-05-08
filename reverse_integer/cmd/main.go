package main

import (
	"fmt"
	"reverse_integer/source"
)

func main() {
	result := source.Reverse(123)
	fmt.Println(result) // 321

	result = source.Reverse(-1005)
	fmt.Println(result) //-5001

	result = source.Reverse(10000212121212121)
	fmt.Println(result) // 0
}
