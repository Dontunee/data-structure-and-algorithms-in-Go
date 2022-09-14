package algoexpert

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			"valid input",
			"abcdcba",
			true,
		},
		{
			"invalid input",
			"shdhds",
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if received := IsPalindrome(test.str); received != test.want {
				t.Error("failed test")
			}
		})
	}
}
