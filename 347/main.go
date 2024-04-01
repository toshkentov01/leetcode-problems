package main

func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]int)

	for i := range nums {
		numsMap[nums[i]]++
	}

	result := make([]int, 0, k)

	var (
		frequent, frequentKey int
	)

	for i := 0; i < k; i++ {
		for key, value := range numsMap {
			if frequent <= value {
				frequent = value
				frequentKey = key
			}
		}

		result = append(result, frequentKey)
		delete(numsMap, frequentKey)

		frequent, frequentKey = 0, 0
	}

	return result
}
