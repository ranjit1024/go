package main

import "fmt"

type vehicle interface {
	wheel() int
}

type car struct {
	move bool
}

type bike struct {
	move bool
}

func (c car) wheel() int {
	return 4
}

func (b bike) wheel() int {
	return 2
}

func wheels(v vehicle) {
	fmt.Printf("%T has %d wheels\n", v, v.wheel())
}

func demoInterfaces() {
	wheels(car{move: true})
	wheels(bike{move: true})
}
