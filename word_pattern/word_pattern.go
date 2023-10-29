package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	sSlice := strings.Split(s, " ")

	if len(sSlice) != len(pattern) {
		return false
	}

	pMap := make(map[byte]string)
	p2Map := make(map[string]byte)

	for i := range pattern {
		value, ok := pMap[pattern[i]]
		if !ok {
			pMap[pattern[i]] = sSlice[i]
		}

		if ok && value != sSlice[i] {
			return false
		}

		value2, ok2 := p2Map[sSlice[i]]
		if !ok2 {
			p2Map[sSlice[i]] = pattern[i]
		}

		if ok2 && value2 != pattern[i] {
			return false
		}

	}

	return true
}
