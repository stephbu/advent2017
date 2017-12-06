package main

import (
	"log"
	"strconv"
)

func main() {

	log.Println(sumOf(validDigits("1122")))
	log.Println(sumOf(validDigits("1111")))
	log.Println(sumOf(validDigits("1234")))
	log.Println(sumOf(validDigits("91212129")))
}

func validDigits(input string) []int {

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

	return digits
}

func sumOf(digits []int) int {

	var total int
	for _, v := range digits {
		total += v
	}

	return total
}
