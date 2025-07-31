package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		password string
		expected ValidationResult
	}{
		{"moreThan8Characters", ValidationResult{true, ""}},
		{"L3ss", ValidationResult{false, "Password must be at least 8 characters"}},
	}

	for _, parameters := range tests {
		actualResult := validate(parameters.password)

		if actualResult != parameters.expected {
			t.Errorf("got %+v wanted %+v", actualResult, parameters.expected)
		}
	}
}
