package main

import (
	"fmt"
	"unicode/utf8"
)

func stringsRunes() {
	const s = "hello"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println("%x", s[i])

	}

	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeLastRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)

	}
}
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'h' {
		fmt.Println("found h")
	}
}
