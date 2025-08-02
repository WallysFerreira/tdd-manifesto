package main

import (
	"slices"
	"testing"
)

func TestMain(t *testing.T) {
	parameters := []struct {
		searchText     string
		expectedResult []string
	}{
		{"", []string{""}},
		{"a", []string{""}},
	}

	for _, parameter := range parameters {
		actualResult := search(parameter.searchText)

		if !slices.Equal(actualResult, parameter.expectedResult) {
			t.Errorf("got %s wanted %s", actualResult, parameter.expectedResult)
		}
	}
}
