package source

func firstUniqChar(s string) int {
	sMap := make(map[byte]int)

	for i := range s {
		sMap[s[i]] += 1
	}

	for i := range s {
		if count, ok := sMap[s[i]]; ok && count == 1 {
			return i
		}
	}

	return -1
}
