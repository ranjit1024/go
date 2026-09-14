package main

import "fmt"

type geometry interface {
	area() int
}
type rect struct {
	height, widht int
}

func (r rect) area() int {
	return r.height * r.widht
}
func main() {
	sqare := rect{
		height: 12,
		widht:  12,
	}

	var g geometry = sqare
	fmt.Println(g.area())

	fmt.Println("Data is the oil")
}
