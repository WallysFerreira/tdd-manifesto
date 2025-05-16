package main

import (
	"regexp"
	"strconv"
	"strings"
)

var allowedSeparators = [2]string{",", "\\n"}

func add(numbersString string) int {
	if numbersString == "" {
		return 0
	}

	return doAdd(numbersString)
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
