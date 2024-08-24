package main

func removeDuplicates(nums []int) int {
	var (
		j         int
		uniqueMap = make(map[int]struct{})
	)

	for i := range nums {
		if _, ok := uniqueMap[nums[i]]; ok {
			continue
		}

		uniqueMap[nums[i]] = struct{}{}
		nums[j] = nums[i]
		j++
	}

	return len(uniqueMap)
}
