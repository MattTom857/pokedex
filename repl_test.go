package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "single",
			expected: []string{"single"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	failure := 0
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Not the right number of words")
			failure = failure + 1
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s is not expected %s", word, expectedWord)
				failure = failure + 1
			}
		}
	}
	success := len(cases) - failure
	fmt.Printf("%d successes", success)
	fmt.Printf("%d failures", failure)
}
