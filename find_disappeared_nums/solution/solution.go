package solution

func findDisappearedNumbers(nums []int) []int {
	numMap := make(map[int]bool)
	result := []int{}

	for i := range nums {
		numMap[nums[i]] = true
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := numMap[i]; !ok {
			result = append(result, i)
		}
	}

	return result
}
