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
		{"", []string{}},
		{"a", []string{}},
		{"Va", []string{"Vancouver", "Valencia"}},
		{"va", []string{"Vancouver", "Valencia"}},
	}

	for _, parameter := range parameters {
		actualResult := search(parameter.searchText)

		for _, eachExpectedResult := range parameter.expectedResult {
			if !slices.Contains(actualResult, eachExpectedResult) {
				t.Errorf("got %s wanted %s", actualResult, parameter.expectedResult)
				return
			}
		}
	}
}
