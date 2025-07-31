package main

import "strings"

func validate(password string) ValidationResult {
	validators := []ValidatorFunc{
		validateLength,
		validateAmountOfNumbers,
		validateCapitalLetters,
		validateSpecialCharacters,
	}

	validationChannel := make(chan ValidationResult, len(validators))

	for _, validator := range validators {
		go func(v ValidatorFunc) {
			result := v(password)
			validationChannel <- result
		}(validator)
	}

	accumulatedErrorMessages := []string{}

	for range validators {
		result := <-validationChannel
		if !result.IsValid {
			accumulatedErrorMessages = append(accumulatedErrorMessages, result.Error)
		}
	}

	if len(accumulatedErrorMessages) != 0 {
		return ValidationResult{false, strings.Join(accumulatedErrorMessages, "\n")}
	}

	return valid()
}
