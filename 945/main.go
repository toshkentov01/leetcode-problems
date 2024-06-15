package main

import "sort"

func minIncrementForUnique(nums []int) int {
	var (
		result int
	)

	sort.Ints(nums)

	increment := 0

	for i := range nums {
		if increment < nums[i] {
			increment = nums[i]
		}

		result += increment - nums[i]
		increment++
	}

	return result
}
