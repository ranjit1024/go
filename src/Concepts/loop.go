package main
import "fmt"

func demoLoops() {
	for i := range 3 {
		fmt.Println(i);
		i++
	}
}
