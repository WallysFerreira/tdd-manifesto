package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"1", 1},
		{"1,2", 3},
	}

	for _, testParameters := range tests {
		testName := fmt.Sprintf("Should return %d when input is %s", testParameters.expected, testParameters.input)

		t.Run(testName, func(t *testing.T) {
			actual, err := add(testParameters.input)

			shouldNotHaveErrors(err, t)
			if actual != testParameters.expected {
				t.Errorf("got %d wanted %d", actual, testParameters.expected)
			}
		})
	}

	t.Run("Should work for multiple numbers", func(t *testing.T) {
		input := "10,30,49"
		expected := 89

		actual, err := add(input)

		shouldNotHaveErrors(err, t)
		if actual != expected {
			t.Errorf("got %d wanted %d", actual, expected)
		}
	})

	t.Run("Should accept newlines as separator", func(t *testing.T) {
		input := "10,30\n49,20"
		expected := 109

		actual, err := add(input)

		shouldNotHaveErrors(err, t)
		if actual != expected {
			t.Errorf("got %d wanted %d", actual, expected)
		}
	})

	t.Run("Should not accept separator at the end", func(t *testing.T) {
		input := "10,20,"

		_, err := add(input)

		if err == nil {
			t.Error("Expected an error due to separator at the end of string")
		}
	})

}

func shouldNotHaveErrors(err error, t *testing.T) {
	if err != nil {
		t.Errorf("unexpected error \"%s\"", err)
	}
}
