package main

import "strconv"

func fizzBuzz(number int) string {
	if number%3 == 0 {
		return "Fizz"
	} else if number%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(number)
	}
}
