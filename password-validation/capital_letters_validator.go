package main

import "unicode"

func validateCapitalLetters(password string) ValidationResult {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return valid()
		}
	}

	return invalidBecause(PasswordNeedsCapitalLetters)
}
