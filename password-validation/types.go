package main

type ValidationResult struct {
	IsValid bool
	Error   string
}

func valid() ValidationResult {
	return ValidationResult{true, ""}
}

func invalidBecause(errorMessage string) ValidationResult {
	return ValidationResult{false, errorMessage}
}

type ValidatorFunc func(password string) ValidationResult
