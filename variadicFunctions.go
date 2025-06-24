package main

import (
	"fmt"
	"math"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func maxium(nums ...float64) float64 {
	max := math.Inf(-1)

}
