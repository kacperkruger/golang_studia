package main

import (
	"flag"
	"fmt"
)

func main() {
	number1 := flag.Int("number1", 0, "first number of sum")
	number2 := flag.Int("number2", 0, "first number of sum")
	flag.Parse()

	if *number1 == 0 {
		fmt.Print("Type first number: ")
		fmt.Scanf("%d\n", number1)

	}

	if *number2 == 0 {
		fmt.Print("Type second number: ")
		fmt.Scanf("%d\n", number2)
	}

	fmt.Println("Sum of numbers", *number1, "+", *number2, "=", *number1+*number2)
}
