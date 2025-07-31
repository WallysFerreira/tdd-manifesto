package main

import "strings"

func validate(password string) ValidationResult {
	validators := []ValidatorFunc{
		validateLength,
		validateAmountOfNumbers,
	}

	accumulatedErrorMessages := []string{}

	for _, validator := range validators {
		if result := validator(password); !result.IsValid {
			accumulatedErrorMessages = append(accumulatedErrorMessages, result.Error)
		}
	}

	if (len(accumulatedErrorMessages) != 0) {
		return ValidationResult{false, strings.Join(accumulatedErrorMessages, "\n")}
	}

	return ValidationResult{true, ""}
}
