package main

func validateLength(password string) ValidationResult {
	if len(password) < 8 {
		return ValidationResult{false, PasswordTooShort}
	}

	return ValidationResult{true, ""}
}
