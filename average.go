package main

import "fmt"

func Average() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number

	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
