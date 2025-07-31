package main

func validate(password string) ValidationResult {
	if len(password) < 8 {
		return ValidationResult{false, "Password must be at least 8 characters"}
	}

	return ValidationResult{true, ""}
}
