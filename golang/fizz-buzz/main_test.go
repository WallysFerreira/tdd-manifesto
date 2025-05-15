package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{31, "31"},
		{30, "Fizz"},
	}

	for _, testParameters := range tests {
		actualString := fizzBuzz(testParameters.number)

		if actualString != testParameters.expected {
			t.Errorf("got %q wanted %q", actualString, testParameters.expected)
		}
	}
}
