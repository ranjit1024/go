package main

import "fmt"

type person struct {
	name    string
	roll_no int
}

func main() {
	fmt.Println("Data is the king")
	p1 := person{name: "Ranjit", roll_no: 12}
	fmt.Println(p1)

	fmt.Println(&person{name: "Samriddha", roll_no: 12})

}
