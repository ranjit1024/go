package main

import "fmt"

type person1 struct {
	name string
	age  int
}

type Employee struct {
	roal string
	person1
}

func main() {
	fmt.Println("Data is the king")
	ranjit := person1{
		name: "Ranjit",
		age:  12,
	}
	fmt.Println(ranjit.name)
}
