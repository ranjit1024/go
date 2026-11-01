package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:age`
}

func demo_json() {
	var user user
	data := []byte(`{"name":"Ranjit", "age":23}`)
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
