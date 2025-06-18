package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func processNumbers(nums []int, callback func(int) int) {
	for i, n := range nums {
		nums[i] = callback(n)
	}
}

func closure() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	nums := []int{1, 2, 3}
	double := func(n int) int { return n * 2 }
	processNumbers(nums, double)
	fmt.Println(nums) //246
}
