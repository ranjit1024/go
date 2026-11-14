package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello World...")
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fmt.Println(fooCmd)
}
