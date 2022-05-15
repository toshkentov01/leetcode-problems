package main

import (
	"subsequence/source"
	"fmt"
)

func main() {
	res := source.IsSubsequence("abc", "ajkfjsbsjfiec")
	fmt.Println(res) // true
}