package main

import "testing"

func TestShouldReturnNumberAsString(t *testing.T) {
	expectedString := "30"

	actualString := fizzBuzz(30)

	if actualString != expectedString {
		t.Errorf("got %q wanted %q", actualString, expectedString)
	}
}
