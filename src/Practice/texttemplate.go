package main

import (
	"os"
	"text/template"
)

func text_template() {
	tmpl := `Hello {{.}}! Welocome to Go`
	t := template.Must(template.New("welcome").Parse(tmpl))
	t.Execute(os.Stdout, "Ranjit")
}
