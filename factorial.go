package main

import "fmt"

func fact(number int64) int64 {

	if number == 1 {
		return 1
	}

	factorialOfNumber := number * fact(number-1)
	return factorialOfNumber

}

func factr() {
	var number int64

	fmt.Scanf("%d\n", &number)
	fmt.Println("Factorial of the  number is ", fact(number))
}
