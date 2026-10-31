package main

import (
	"fmt"
	"os"
	"text/template"
)

func text_template() {
	tmp1 := `Hello {{.}} Welcome to go`
	t := template.Must(template.New("Welcome").Parse(tmp1))
	t.Execute(os.Stdout, "Ranjit")
	fmt.Println("Hello World...")
}
