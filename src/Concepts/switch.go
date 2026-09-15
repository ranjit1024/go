package main
import "fmt"

func demoSwitch() {
	i := 2;
	fmt.Println("Write", i , " as ");
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	}
	fmt.Println("Hello World");
	
}