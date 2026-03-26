package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {
	case time.Sunday:
		fmt.Println("Today is sunday")
	default:
		fmt.Println("Today is weekday")
	}

}
