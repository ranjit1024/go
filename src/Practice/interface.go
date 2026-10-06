package main

import "fmt"

type verical interface {
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

func wheels(v verical) {
	fmt.Println(v)
	fmt.Println(v.wheel())
}

func interface_demo() {
	b := bike{
		move: true,
	}
	wheels(b)
	fmt.Println("Understanding Interface")
}
