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

func demoEmbedding() {
	c := container{
		base: base{num: 1},
		str:  "Ranjit",
	}
	fmt.Println(c)
	fmt.Println(c.describe())
}
