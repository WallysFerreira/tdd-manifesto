package main

import (
	"strconv"
	"strings"
)

func add(numbersString string) int {
	if numbersString == "" {
		return 0
	}

	return doAdd(numbersString)
}

func doAdd(numbersString string) int {
	numbers := strings.Split(numbersString, ",")

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
