package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var allowedSeparators = [2]string{",", "\\n"}
var chosenSeparator string

func add(numbersString string) (int, error) {
	if numbersString == "" {
		return 0, nil
	} else if lastCharacterIsSeparator(numbersString) {
		return 0, errors.New("separator at the end of String not allowed")
	}

	setChosenSeparator(numbersString)

	return doAdd(numbersString), nil
}

func setChosenSeparator(numbersString string) {
	if (len(numbersString) > 2 && numbersString[0:2] == "//") && (strings.Contains(numbersString, "\n")) {
		customSeparator := strings.SplitN(numbersString, "\n", 2)[0][2:]
		chosenSeparator = customSeparator
	}
}

func lastCharacterIsSeparator(numbersString string) bool {
	lastCharacter := string(numbersString[len(numbersString)-1])

	return anySeparatorMatch(lastCharacter)
}

func anySeparatorMatch(character string) bool {
	for _, separator := range allowedSeparators {
		if character == separator {
			return true
		}
	}

	return false
}

func doAdd(numbersString string) int {
	regex := regexp.MustCompile(makeSeparatorsRegex())
	numbers := getNumbersList(numbersString, regex)

	if len(numbers) == 1 {
		number, _ := strconv.Atoi(numbers[0])
		return number
	}

	return sum(numbers)
}

func getNumbersList(numbersString string, regex *regexp.Regexp) []string {
	if chosenSeparator != "" {
		numbersPart := strings.SplitN(numbersString, "\n", 2)[1]
		return regex.Split(numbersPart, -1)
	} else {
		return regex.Split(numbersString, -1)
	}
}

func sum(numbers []string) int {
	total := 0

	for _, numberAsString := range numbers {
		number, _ := strconv.Atoi(numberAsString)
		total += number
	}

	return total
}

func makeSeparatorsRegex() string {
	if chosenSeparator != "" {
		return `[` + regexp.QuoteMeta(chosenSeparator) + `]`
	}

	return `[` + strings.Join(allowedSeparators[:], "") + `]`
}
