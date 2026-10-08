package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with %v", b.num)
}

type container struct {
	base
	str string
}

func demo_struct_embeding() {
	co := container{
		base: base{
			num: 1,
		},
		str: "Sammmmu",
	}
	fmt.Println(co)
	fmt.Println("hello World....")
}
