package main

func sortedSquares(nums []int) []int {
	var (
		left  int
		right = len(nums) - 1
	)

	result := make([]int, len(nums))
	resultIndex := len(result) - 1

	for {
		if left == right {
			result[resultIndex] = nums[right] * nums[right]

			break
		}

		if left > right {
			break
		}

		leftSquare := square(nums[left])
		rightSquare := square(nums[right])

		if leftSquare < rightSquare {
			result[resultIndex] = rightSquare
			resultIndex--
			right--

			continue
		}

		if leftSquare > rightSquare {
			result[resultIndex] = leftSquare
			resultIndex--
			left++

			continue
		}

		if leftSquare == rightSquare {
			result[resultIndex] = rightSquare
			resultIndex--
			right--

			result[resultIndex] = leftSquare
			resultIndex--
			left++
		}
	}

	return result
}

func square(a int) int {
	return a * a
}
