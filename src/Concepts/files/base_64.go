package main

import (
	"encoding/base64"
	"fmt"
)

func base_demo() {
	data := "abc123!?$*&()'-=@~"
	encode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encode)
	fmt.Println("Hello World...")
}
