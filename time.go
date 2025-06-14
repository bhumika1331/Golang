package main

import (
	"fmt"
	"time"
)

func timeFunction() {

	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
}
