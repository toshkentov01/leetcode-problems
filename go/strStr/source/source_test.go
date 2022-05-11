package source

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

type TestCase struct {
	Haystack string
	Needle   string
	Expected int
}

func TestStrStr(t *testing.T) {
	testCases := []TestCase{
		{
			Haystack: "hello",
			Needle:   "ll",
			Expected: 2,
		},
		{
			Haystack: "apple",
			Needle:   "",
			Expected: 0,
		},
		{
			Haystack: "hello",
			Needle:   "b",
			Expected: -1,
		},
	}

	t.Run("Testing StrStr method", func(t *testing.T) {
		for _,  testCase := range testCases {
			result := StrStr(testCase.Haystack, testCase.Needle)
			
			if difference := cmp.Diff(testCase.Expected, result); difference != "" {
				t.Error(difference)
			}
		}
	})
}
