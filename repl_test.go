package main

import "testing"


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "         ",
			expected: []string{},
		},
		{
			input: "Golang",
			expected: []string{"Golang"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Input: %q. Length mismatch. Got: %d, Expected: %d",
				c.input, len(actual), len(c.expected))
			continue // Skip further checks for this case
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Input: %q. Word mismatch at index %d. Got: %q, Expected: %q",
					c.input, i, word, expectedWord)
			}
		}
	}
}