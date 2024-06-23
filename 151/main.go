package main

import "strings"

func reverseWords(s string) string {
	slice := strings.Fields(s)
	str := strings.Builder{}

	for i := len(slice) - 1; i >= 0; i-- {
		str.WriteString(slice[i])

		if i != 0 {
			str.WriteString(" ")
		}
	}

	return str.String()
}
