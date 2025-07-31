package main

type ValidationResult struct {
	IsValid bool
	Error   string
}

type ValidatorFunc func(password string) ValidationResult
