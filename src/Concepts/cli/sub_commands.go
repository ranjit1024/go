package main

import (
	"flag"
	"fmt"
	"os"
)

func subcommand() {
	fmt.Println("Hello World...")
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnabled := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'foo'  or 'bar' subcommands")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[:2])
		fmt.Println("Subcommand 'foo' ")
		fmt.Println("	enable : ", *fooEnabled)
		fmt.Println("	Name : ", *fooName)
		fmt.Println("	tail:	", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("sumcommand 'bar' ")
		fmt.Println("	Level:	", *barLevel)
		fmt.Println("	Tail:	", barCmd.Args())
	}

	fmt.Println(fooCmd)
}
