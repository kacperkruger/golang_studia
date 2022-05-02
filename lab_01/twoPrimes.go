package main

import (
	"fmt"
	"math"
)

type booleanFunction func(int) bool

func isPrime(number int) bool {
	return isPrimeHelper(number, 2)
}

func isPrimeHelper(number int, divider int) bool {
	if number < 2 {
		return false
	}
	if number == 2 {
		return true
	}
	if number%divider == 0 {
		return false
	}
	if math.Sqrt(float64(number)) < float64(divider) {
		return true
	}
	return isPrimeHelper(number, divider+1)
}

func filter(arr []int, fn booleanFunction) []int {
	var result []int
	for i := range arr {
		if fn(arr[i]) == true {
			result = append(result, arr[i])
		}
	}
	return result
}

func makeRange(max int) []int {
	array := make([]int, max+1)
	for i := range array {
		array[i] = i
	}
	return array
}

func sieveOfEratosthenes(number int) []int {
	array := makeRange(number)
	filteredArray := filter(array, isPrime)
	return filteredArray
}

func main() {
	var number int
	fmt.Print("Type number on which sieve must stop: ")
	fmt.Scan(&number)
	fmt.Println(sieveOfEratosthenes(number))
}
