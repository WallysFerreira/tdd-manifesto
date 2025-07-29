package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var allowedSeparators = [2]string{",", "\\n"}

func add(numbersString string) (int, error) {
	if numbersString == "" {
		return 0, nil
	} else if lastCharacterIsSeparator(numbersString) {
		return 0, errors.New("separator at the end of String not allowed")
	}

	return doAdd(numbersString), nil
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
	numbers := regex.Split(numbersString, -1)

	if len(numbers) == 1 {
		number, _ := strconv.Atoi(numbers[0])
		return number
	}

	return sum(numbers)
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
	return `[` + strings.Join(allowedSeparators[:], "") + `]`
}
