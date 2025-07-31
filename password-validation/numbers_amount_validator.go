package main

import "unicode"

func validateAmountOfNumbers(password string) ValidationResult {
	if amountOfNumbersIn(password) < 2 {
		return ValidationResult{false, PasswordNeedsMoreNumbers}
	}

	return valid()
}

func amountOfNumbersIn(password string) int {
	count := 0

	for _, char := range password {
		if unicode.IsDigit(char) {
			count++
		}
	}

	return count
}
