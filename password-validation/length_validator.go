package main

func validateLength(password string) ValidationResult {
	if len(password) < 8 {
		return invalidBecause(PasswordTooShort)
	}

	return valid()
}
