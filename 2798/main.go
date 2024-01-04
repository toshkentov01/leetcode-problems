package main

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var count int

	for i := range hours {
		if hours[i] >= target {
			count++
		}
	}

	return count
}
