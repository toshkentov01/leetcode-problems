package source

import (
	"fmt"
	"testing"
	"time"
)

func TestIsSubsequence(t *testing.T) {
	testCases := []TestCase{
		{
			firstString:  "abc",
			secondString: "alkldsbklklsc",
			expected:     true,
		},
		{
			firstString:  "abf",
			secondString: "aklklklflklb",
			expected:     false,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing IsSubsecquence function at: %s", time.Now()), func(t *testing.T) {
			result := IsSubsequence(testCase.firstString, testCase.secondString)

			if result != testCase.expected {
				t.Errorf("Error while tesing IsSubsecquence method. Want %v, got %v", testCase.expected, result)
			}
		})
	}
}
