package main

import "fmt"

func main() {

	var a, b int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	SumOfTwoInteger(a, b)

}

func SumOfTwoInteger(a int, b int) int {
	c := a + b

	fmt.Println("The Sum of 2 Integers", c)
	return c
}
