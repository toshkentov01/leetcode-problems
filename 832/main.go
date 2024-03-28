package main

func flipAndInvertImage(image [][]int) [][]int {
	result := make([][]int, 0, len(image))

	for i := range image {
		in := make([]int, 0, len(image))
		for j := len(image[i]) - 1; j >= 0; j-- {
			in = append(in, invert(image[i][j]))
		}

		result = append(result, in)
	}

	return result
}

func invert(input int) int {
	if input == 1 {
		return 0
	}

	return 1
}
