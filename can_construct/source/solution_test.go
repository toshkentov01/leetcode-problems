package source

import "testing"

type TestCase struct {
	RansomNote string
	Magazine   string
	Expected   bool
}

func TestCanConstruct(t *testing.T) {
	testCases := []TestCase{
		{
			RansomNote: "a",
			Magazine:   "b",
			Expected:   false,
		},
		{
			RansomNote: "abc",
			Magazine:   "bba",
			Expected:   false,
		},
		{
			RansomNote: "ffg",
			Magazine:   "ffhg",
			Expected:   true,
		},
	}

	for _, data := range testCases {
		result := CanConstruct(data.RansomNote, data.Magazine)
		
		if result != data.Expected {
			t.Errorf("FAILED: WANT: %v, GOT: %v", data.Expected, result)	
		}
	}
}
