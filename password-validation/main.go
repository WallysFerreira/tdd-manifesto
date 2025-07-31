package main

func validate(password string) ValidationResult {
	if failsLengthCheck(password) {
		return ValidationResult{false, PasswordTooShort}
	}

	return ValidationResult{true, ""}
}

func failsLengthCheck(password string) bool {
	return len(password) < 8
}
