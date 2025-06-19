package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func pointers() {

	i := 1
	fmt.Println("inital:", i)

	zeroVal(i)
	fmt.Println("Zeroval:", i)

	zeroptr(&i)
	fmt.Println("Zeroptr:", i)

	fmt.Println("pointer", &i)

}
