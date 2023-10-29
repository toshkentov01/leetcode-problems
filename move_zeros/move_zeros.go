package movezeros

func moveZeroes(nums []int) {
	var (
		position int
	)

	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[position] = nums[position], nums[i]
			position++
		}
	}
}
