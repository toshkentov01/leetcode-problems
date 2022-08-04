package source

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max, min := 0, prices[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]

		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max
}
