package main

import "testing"

func TestMain(t *testing.T) {
	parameters := []struct {
		searchText     string
		expectedResult string
	}{
		{"", ""},
		{"a", ""},
	}

	for _, parameter := range parameters {
		actualResult := search(parameter.searchText)

		if actualResult != parameter.expectedResult {
			t.Errorf("got %s wanted %s", actualResult, parameter.expectedResult)
		}
	}
}
