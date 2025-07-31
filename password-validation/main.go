package main

func validate(password string) ValidationResult {
	validators := []ValidatorFunc{
		validateLength,
		validateAmountOfNumbers,
	}

	for _, validator := range validators {
		if result := validator(password); !result.IsValid {
			return result
		}
	}

	return ValidationResult{true, ""}
}
