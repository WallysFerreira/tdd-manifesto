package main

import (
	"regexp"
	"tdd/search/city"
)

func search(searchText string) (cityNamesMatched []string) {
	if searchText == "*" {
		return allCitiesNames()
	}

	if len(searchText) < 2 {
		return
	}

	for _, city := range city.Cities() {
		regexPattern, _ := regexp.Compile(caseInsensitive(searchText))
		matches := regexPattern.Match([]byte(city.Name))
		if matches {
			cityNamesMatched = append(cityNamesMatched, city.Name)
		}
	}
	return
}

func caseInsensitive(text string) string {
	return "(?i)" + text
}

func allCitiesNames() (names []string) {
	for _, city := range city.Cities() {
		names = append(names, city.Name)
	}

	return
}
