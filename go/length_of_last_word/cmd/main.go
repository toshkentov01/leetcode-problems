package main

import (
	"fmt"
	"lenght/source"
)

func main() {
	result := source.LengthOfLastWord("Hello World")
	fmt.Println(result) // 5

	result2 := source.LengthOfLastWord("My name is Sardor   ")
	fmt.Println(result2)
}
