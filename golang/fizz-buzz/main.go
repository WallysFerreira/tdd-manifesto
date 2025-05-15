package main

import "strconv"

func fizzBuzz(number int) string {
	if number%3 == 0 {
		return "Fizz"
	} else {
		return strconv.Itoa(number)
	}
}
