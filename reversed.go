package main

import "fmt"

func main() {

	var k string

	fmt.Scanf("%s", &k)

	fmt.Println(ReverseString(k))

}

func ReverseString(c string) string {

	runes := []rune(c)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
