package main

import (
	"fmt"
	"text/template"
)

func text_template() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Valeu is {{.}}\n")
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello World...")
}
