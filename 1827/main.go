package main

func minOperations(nums []int) int {
	var (
		result int
		last   int
	)

	for i := range nums {
		if last < nums[i] {
			last = nums[i]

			continue
		}

		result += (last + 1) - nums[i]
		last++
	}

	return result
}
