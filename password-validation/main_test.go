package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		password string
		expected ValidationResult
	}{
		{"moreThan8Characters2", ValidationResult{true, ""}},
		{"L3ss2", ValidationResult{false, PasswordTooShort}},
		{"Only1Number", ValidationResult{false, PasswordNeedsMoreNumbers}},
	}

	for _, parameters := range tests {
		actualResult := validate(parameters.password)

		if actualResult != parameters.expected {
			t.Errorf("got %+v wanted %+v", actualResult, parameters.expected)
		}
	}
}
