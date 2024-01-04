package main

import "slices"

func numberGame(nums []int) []int {
	arr := make([]int, len(nums))

	slices.Sort(nums)

	for i := 0; i < len(nums)-1; i += 2 {
		a, b := nums[i], nums[i+1]

		arr[i], arr[i+1] = b, a
	}

	return arr
}
