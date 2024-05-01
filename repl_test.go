package main

import "testing"

func TestCleanUserInput(test *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Yo Mama Ugly  ",
			expected: []string{"yo", "mama", "ugly"},
		},
	}

	for _, testCase := range testCases {
		actual := cleanUserInput(testCase.input)
		if len(actual) != len(testCase.expected) {
			test.Errorf("lengths don't match: '%v' vs '%v'", actual, testCase.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := testCase.expected[i]
			if word != expectedWord {
				test.Errorf("cleanInput(%v) == %v, expected %v", testCase.input, actual, testCase.expected)
			}
		}
	}
}
