package main

import "strconv"

func fizzBuzz(number int) string {
	if isDivisableBy(number, 3) && isDivisableBy(number, 5) {
		return "FizzBuzz"
	} else if isDivisableBy(number, 3) {
		return "Fizz"
	} else if isDivisableBy(number, 5) {
		return "Buzz"
	} else {
		return strconv.Itoa(number)
	}
}

func isDivisableBy(number int, divisor int) bool {
	return number%divisor == 0
}
