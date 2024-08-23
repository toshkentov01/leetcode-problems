package main

func removeElement(nums []int, val int) int {
	var (
		counter, i int
		lastIndex  = len(nums) - 1

		idf = len(nums)
	)

	for i := range nums {
		if nums[i] != val {
			counter++
		}
	}

	for i != idf {
		if nums[i] == val && nums[lastIndex] != val {
			nums[i], nums[lastIndex] = nums[lastIndex], nums[i]

			lastIndex--
			i++
			idf--

			continue
		}

		if nums[i] == val && nums[lastIndex] == val {
			lastIndex--
			idf--

			continue
		}

		i++
	}

	return counter
}
