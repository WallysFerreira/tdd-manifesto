package main

import (
	"regexp"
	"tdd/search/city"
)

func search(searchText string) (cityNamesMatched []string) {
	if len(searchText) < 2 {
		return
	}

	for _, city := range city.Cities() {
		regexPattern, _ := regexp.Compile(searchText)
		matches := regexPattern.Match([]byte(city.Name))
		if matches {
			cityNamesMatched = append(cityNamesMatched, city.Name)
		}
	}
	return
}
