package main

import (
	"flag"
	"fmt"
)

func flags_demon() {
	fmt.Println("fadfa")
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word: ", *wordPtr)
	fmt.Println("Numb: ", *numPtr)
	fmt.Println("Fork: ", *forkPtr)
}
