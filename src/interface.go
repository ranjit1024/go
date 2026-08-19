package main

import "fmt"

type geometry interface {
	area() float32
	para() float32
}
type rect struct {
	width, height int
}

func (r rect) area() float32 {
	return float32(r.width * r.height)
}

func (r rect) para() float32 {
	return float32(r.height + r.width)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.para())
}

func main() {

	fmt.Println("DAta is the king")
}
