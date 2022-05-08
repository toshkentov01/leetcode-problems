package main

import (
	"find_the_difference/source"
	"fmt"
)

func main() {
	result := source.FindTheDifference("abcd", "abcde")
	fmt.Println("First Result: ", string(result)) // e

	result = source.FindTheDifference("aa", "aab")
	fmt.Println("Second Result: ", string(result)) // a
}
