package source

import (
	"fmt"
	"testing"
	"time"
)

type TestCase struct {
	Words1   []string
	Words2   []string
	Expected int
}

var testCases []TestCase = []TestCase{
	{
		Words1:   []string{"leetcode", "hackerrank", "apple", "banana"},
		Words2:   []string{"leetcode", "hackerrank", "apple"},
		Expected: 3,
	},
	{
		Words1:   []string{"Ubuntu", "Manjaro", "Kubuntu", "Ubuntu"},
		Words2:   []string{"Mint", "Ubuntu", "Manjaro"},
		Expected: 1,
	},
}

func TestCountWords(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing CountWords function at %s", time.Now()), func(t *testing.T) {
			result := CountWords(testCase.Words1, testCase.Words2)

			if result != testCase.Expected {
				t.Errorf("Test FAILED. WANTED: %v, GOT: %v", testCase.Expected, result)
			}
		})
	}
}
