package main

import (
	"fmt"
	"os"
	"strings"
)

func demo_env() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO", os.Getenv("FOO"))
	fmt.Print("fasdfadf")

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
