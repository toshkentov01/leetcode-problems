package main

func findMin(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		mid := (start + end) / 2

		if nums[mid] < nums[end] {
			// The minimum is in the left half, including mid
			end = mid
		} else {
			// The minimum is in the right half, excluding mid
			start = mid + 1
		}
	}

	// At the end of the loop, start == end, which is the index of the minimum
	return nums[start]
}
