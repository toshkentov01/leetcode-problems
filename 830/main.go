package main

func largeGroupPositions(s string) [][]int {
	var (
		start, end int
		count      int = 1
	)

	result := make([][]int, 0, len(s))

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			if count == 1 {
				start = i
			}

			count++
			end = i + 1

			if i == len(s)-2 && count >= 3 {
				result = append(result, []int{start, end})
			}

			continue
		}

		if count >= 3 {
			result = append(result, []int{start, end})
		}

		count = 1
		end = 0
		start = i
	}

	return result
}
