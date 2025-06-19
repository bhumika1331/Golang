package main

import (
	"fmt"
	"reflect"
)

func learningSlices() {
	s := []string{"Hello", "hi", "bye", "Tata"}

	fmt.Println("Length of Slice", len(s))

	fmt.Println("Reflecting the type", reflect.TypeOf(s))

	c := s[:2]
	fmt.Println(c)

	d := []string{"okay"}

	fmt.Println(d)

	s = append(s, d...)
	fmt.Println(s)
}
