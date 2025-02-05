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
			input:    "HELLO  worlD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) == 0 && len(c.expected) != 0 {
			t.Errorf("given %v got empty slice when length %v was expected", c.input, len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("given %v want %v got %v", c.input, expectedWord, word)
			}
		}
	}

}
