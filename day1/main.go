package main

import (
	"log"
	"strconv"
)

func main() {

	sumOfSequential("1122")
	sumOfSequential("1111")
	sumOfSequential("1234")
	sumOfSequential("91212129")
}

func sumOfSequential(input string) int {

	var digits []int
	var last int

	for i := range input {
		current, _ := strconv.Atoi(string(input[i]))
		if current == last {
			digits = append(digits, current)
		}
		last = current
	}

	if first, _ := strconv.Atoi(string(input[0])); last == first {
		digits = append(digits, last)
	}

	log.Println(digits, sumOfDigits(digits))

	return 0
}

func sumOfDigits(digits []int) int {

	var total int
	for _, v := range digits {
		total += v
	}

	return total
}
