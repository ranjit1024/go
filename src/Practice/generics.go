package main
import "fmt"

func Index[T comparable](s []T, target T)int{
	for _, value := range s{
		if value == target{
			return 1
		}
	}
	return -1
}

func demo_generics(){
	index := Index([]int{1,2,3,4},2);
	fmt.Println(index)
	fmt.Println("Hello World")
}
