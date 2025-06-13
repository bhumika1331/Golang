package main

import "fmt"

func check() (int, int) {
	return 1, 2
}

func check2() {
	a, b := check()
	fmt.Println(a)
	fmt.Println(b)

	_, c := check()
	fmt.Println(c)
	if c == 0 {
		fmt.Println("c has zero value")
	}
}
