package main

import "fmt"

func main() {
	var k int

	fmt.Scanf("%d", &k)

	if IsTheNoPrime(k) {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}

func IsTheNoPrime(k int) bool {

	if k <= 1 {
		return false
	}

	for i := 2; i*i <= k; i++ {
		if k%i == 0 {
			return false
		}
	}
	return true

}
