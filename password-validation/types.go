package main

type ValidationResult struct {
	IsValid bool
	Error   string
}

func valid() ValidationResult {
	return ValidationResult{true, ""}
}

type ValidatorFunc func(password string) ValidationResult
