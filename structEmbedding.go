package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}


type container struct {
	base
	str string
}

func main() {
	co := container {
		base: base{
			nums: 1;
		},
		str: "some name",
	}

	fmt.Printf("co={nums}")
	}