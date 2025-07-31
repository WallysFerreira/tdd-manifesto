package main

import "unicode"

func validateCapitalLetters(password string) ValidationResult {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return ValidationResult{true, ""}
		}
	}

	return ValidationResult{false, PasswordNeedsCapitalLetters}
}