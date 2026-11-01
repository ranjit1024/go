package main

import (
	"encoding/json"
	"fmt"
)

type reponse1 struct {
	Page   int
	Fruits []string
}
type reposen2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func demo_json() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	fmt.Println("Hello World...")
}
