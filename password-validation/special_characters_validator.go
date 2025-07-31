package main

import "unicode"

func validateSpecialCharacters(password string) ValidationResult {
	for _, char := range password {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			return valid()
		}
	}

	return invalidBecause(PasswordNeedsSpecialCharacters)
}
