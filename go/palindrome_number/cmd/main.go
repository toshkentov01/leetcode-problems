package main

import (
	"fmt"
	"palindrome_number/source"
)

func main() {
	palindrome := source.IsPalindrome(12221)

	if palindrome {
		fmt.Println("Palindrome")
	
	} else {
		fmt.Println("Not Palindrome")
	}
}
