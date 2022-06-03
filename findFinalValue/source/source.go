package source

func FindFinalValue(nums []int, original int) int {
    var found bool
    
    for index, num := range nums {
		if original == num {
			original = FindFinalValue(nums, original *  2)
			found = true
			
		} else {
			found = false
		}

		if index == len(nums) - 1 && !found {
			return original
		}
	}

	return original
}