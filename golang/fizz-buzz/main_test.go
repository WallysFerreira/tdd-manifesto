package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("Should return number as string", func(t *testing.T) {
		expectedString := "31"

		actualString := fizzBuzz(31)

		if actualString != expectedString {
			t.Errorf("got %q wanted %q", actualString, expectedString)
		}
	})

	t.Run("Should return 'Fizz' when number is multiple of three", func(t *testing.T) {
		expectedString := "Fizz"

		actualString := fizzBuzz(30)

		if actualString != expectedString {
			t.Errorf("got %q wanted %q", actualString, expectedString)
		}
	})
}
