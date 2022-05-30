package source

func PercentageLetter(s string, letter byte) int {
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			count += 1
		}
	}
    
	return int(float64(count) / float64(len(s)) * 100)
}