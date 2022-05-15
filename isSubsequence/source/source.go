package source

func IsSubsequence(s string, t string) bool {
	counter := 0

	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(t); i++ {
		if t[i] == s[counter] {
			counter += 1
		}

		if counter == len(s) {
			return true
		}
	}
	
	return false
}