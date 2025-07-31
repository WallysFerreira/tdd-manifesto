package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		password string
		expected ValidationResult
	}{
		{"moreThan8Ch@racters2", ValidationResult{true, ""}},
		{"L3ss2!", ValidationResult{false, PasswordTooShort}},
		{"Only1Number*", ValidationResult{false, PasswordNeedsMoreNumbers}},
		{"noc4p1tall3tters#", ValidationResult{false, PasswordNeedsCapitalLetters}},
		{"N0sp3cialChars", ValidationResult{false, PasswordNeedsSpecialCharacters}},
		{"bad", ValidationResult{false, PasswordTooShort + "\n" + PasswordNeedsMoreNumbers + "\n" + PasswordNeedsCapitalLetters + "\n" + PasswordNeedsSpecialCharacters}},
	}

	for _, parameters := range tests {
		actualResult := validate(parameters.password)

		if actualResult != parameters.expected {
			t.Errorf("got %+v wanted %+v", actualResult, parameters.expected)
		}
	}
}
