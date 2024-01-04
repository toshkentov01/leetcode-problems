package main

func minOperations(nums []int, k int) int {
	numMap := make(map[int]struct{})

	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := numMap[nums[i]]; !ok && nums[i] <= k {
			numMap[nums[i]] = struct{}{}
		}

		if len(numMap) == k {
			return len(nums) - i
		}
	}

	return 0
}
