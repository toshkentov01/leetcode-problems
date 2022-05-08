package source

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if result := digits[i] + 1; result != 10 {
			digits[i] += 1

			return digits
		}

		digits[i] = 0
		
		if i == 0 {
			// If i == 0 and this loop still works (not returned), then it means that
			// all elements of the digits slice are equal to 0
			digits = append(digits, 0)
			digits[0] = 1
		}
	}
	return digits
}