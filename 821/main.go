package main

import "fmt"

func shortestToChar(s string, c byte) []int {
	indexSlice := make([]int, 0, len(s))

	result := make([]int, len(s))

	for i := range s {
		if s[i] == c {
			indexSlice = append(indexSlice, i)
		}
	}

	var (
		counter        int
		lastEqualIndex int
		equalFound     bool
	)

	for i := range s {
		if s[i] != c {
			if !equalFound {
				result[i] = indexSlice[counter] - i
				continue
			}

			distance := indexSlice[counter] - i
			if distance < 0 {
				distance = distance * -1
			}

			if abs(lastEqualIndex, i) >= distance {
				result[i] = distance
				continue
			}

			result[i] = abs(lastEqualIndex, i)
			continue
		}

		lastEqualIndex = i
		equalFound = true
		result[i] = 0

		if counter+1 != len(indexSlice) {
			counter++
		}
	}

	return result
}

func abs(a, b int) int {
	fmt.Println(a, b)
	if a-b < 0 {
		return (a - b) * -1
	}

	return a - b
}
