package main

import (
	"fmt"
	"os"
	"text/template"
)

func demoTextTemplate() {
	text := `Hello {{.}} Welcome to go`
	t := template.Must(template.New("welcome").Parse(text))
	if err := t.Execute(os.Stdout, "Ranjit"); err != nil {
		fmt.Println("template error:", err)
	}
	fmt.Println()
}
