package main

import "unicode"

func validate(password string) ValidationResult {
	if failsLengthCheck(password) {
		return ValidationResult{false, PasswordTooShort}
	}

	if containsLessNumbersThanNeeded(password) {
		return ValidationResult{false, PasswordNeedsMoreNumbers}
	}

	return ValidationResult{true, ""}
}

func failsLengthCheck(password string) bool {
	return len(password) < 8
}

func containsLessNumbersThanNeeded(password string) bool {
	count := 0

	for _, char := range password {
		if unicode.IsDigit(char) {
			count++
		}
	}

	return count < 2
}
